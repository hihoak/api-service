package api

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	custom_errors "hihoak/api-service/internal/pkg/custom-errors"
	spotify "hihoak/api-service/internal/pkg/spotify-structs"
	"hihoak/api-service/internal/pkg/utils/logger"
	pb "hihoak/api-service/pkg/pb/hihoak/api-service"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const goroutinesCount = 4

func (m *musicApi) SpotifySearchTopArtists(ctx context.Context,
	req *pb.SpotifySearchTopArtistsRequest) (*pb.SpotifySearchTopArtistsResponse, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg(fmt.Sprintf("Start SpotifyFindArtist with req %v", req))
	minFollowers := req.GetMinFollowers()
	minPopularity := req.GetMinPopularity()
	count := req.GetCount()
	if count <= 0 {
		return makeSpotifySearchTopArtistResponse(nil), nil
	}

	filteredArtistsByOffset := make(map[int][]*pb.SpotifyArtist)
	var totalArtists int64 = 0
	mu := sync.Mutex{}
	errs, ctx := errgroup.WithContext(ctx)
	for i := 0; i < goroutinesCount; i++ {
		idx := i

		body := url.Values{}
		var query []string
		if len(req.GetGenres()) != 0 {
			query = append(query, fmt.Sprintf("genre:%s", strings.Join(req.GetGenres(), ",")))
		}
		if len(req.GetExcludeGenres()) != 0 {
			for _, genre := range req.GetExcludeGenres() {
				query = append(query, fmt.Sprintf("-genre:%s", genre))
			}
		}
		body.Add("q", strings.Join(query, " "))
		body.Add("type", "artist")
		reqLimit := 50
		body.Add("limit", strconv.Itoa(reqLimit))

		headers := http.Header{}
		token, ok := ctx.Value(TokenKey).(string)
		if !ok {
			return nil, fmt.Errorf("can't get authorization token")
		}
		headers.Set("Authorization", "Bearer " + token)
		headers.Set("Content-Type", "application/json")
		errs.Go(func() error {
			var resp *http.Response
			var localErr error
			defer func() {
				resp.Body.Close()
			}()

			step := 0
			for {
				offset := idx * reqLimit + step * reqLimit * goroutinesCount
				body.Set("offset", strconv.Itoa(offset))
				step += 1
				apiRequest, err := http.NewRequest("GET",
					fmt.Sprintf("https://api.spotify.com/v1/search?%s", body.Encode()), nil)
				if err != nil {
					return fmt.Errorf("can't create request to Spotify API, %w", err)
				}
				apiRequest.Header = headers

				log.Debug().Msg(fmt.Sprintf("Successfully create request to Spotify API. req: %v", apiRequest))
				resp, localErr = http.DefaultClient.Do(apiRequest)
				if localErr != nil {
					return fmt.Errorf("can't do request to Spotify API, %w", localErr)
				}

				data, localErr := io.ReadAll(resp.Body)
				if localErr != nil {
					return fmt.Errorf("%s can't read response body %w. Status code %d",
						ErrorMessage, localErr, resp.StatusCode)
				}
				log.Debug().Msg(fmt.Sprintf("Got response from Spotify API. resp: %v", resp))

				if resp.StatusCode == http.StatusNotFound {
					return nil
				}
				if resp.StatusCode != http.StatusOK {
					jsonResp := &BadSpotifySearchArtist{}
					if localErr = json.Unmarshal(data, jsonResp); localErr != nil {
						return fmt.Errorf("%s can't parse response body %w. Status code %d",
							ErrorMessage, localErr, resp.StatusCode)
					}

					return fmt.Errorf("%s %s. Status code %d", ErrorMessage, jsonResp.Error.Message, resp.StatusCode)
				}

				jsonResp := &OkSpotifySearchResponse{}
				if localErr = json.Unmarshal(data, jsonResp); localErr != nil {
					return fmt.Errorf("%s can't parse response body %w. Status code %d",
						ErrorMessage, localErr, resp.StatusCode)
				}

				// если пришел ответ с пустым списком исполнителей значит такого жанра нет в принципе
				if len(jsonResp.Artists.Items) == 0 {
					return custom_errors.ErrArtistNotFound
				}

				mu.Lock()
				for _, artist := range jsonResp.Artists.Items {
					if artist.Followers.Total > minFollowers && artist.Popularity > minPopularity {
						filteredArtistsByOffset[offset] = append(filteredArtistsByOffset[offset],
							spotify.FromStructToPb(artist))
						totalArtists += 1
					}
				}
				if totalArtists >= count {
					mu.Unlock()
					return nil
				}
				mu.Unlock()
			}
		})
	}
	err := errs.Wait()
	filteredArtists := make([]*pb.SpotifyArtist, 0, count)
	offsets := make([]int, 0, len(filteredArtistsByOffset))
	for offset := range filteredArtistsByOffset {
		offsets = append(offsets, offset)
	}
	sort.Ints(offsets)

	for idx := 0; idx < len(offsets) && len(filteredArtists) < int(count); idx++ {
		for jdx := 0; jdx < len(filteredArtistsByOffset[offsets[idx]]) && len(filteredArtists) < int(count); jdx++ {
			filteredArtists = append(filteredArtists, filteredArtistsByOffset[offsets[idx]][jdx])
		}
	}
	return makeSpotifySearchTopArtistResponse(filteredArtists), err
}

func makeSpotifySearchTopArtistResponse(artists []*pb.SpotifyArtist) *pb.SpotifySearchTopArtistsResponse {
	return &pb.SpotifySearchTopArtistsResponse {
		Artists: artists,
	}
}

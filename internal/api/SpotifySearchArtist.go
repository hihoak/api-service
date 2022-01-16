package api

import (
	"context"
	"encoding/json"
	"fmt"
	custom_errors "hihoak/api-service/internal/pkg/custom-errors"
	spotify "hihoak/api-service/internal/pkg/spotify-structs"
	"hihoak/api-service/internal/pkg/utils/logger"
	pb "hihoak/api-service/pkg/pb/hihoak/api-service"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const ErrorMessage = "bad request to Spotify API. Error message:"

type BadSpotifySearchArtistError struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
}

type BadSpotifySearchArtist struct {
	Error BadSpotifySearchArtistError `json:"error"`
}

type OkSpotifySearchResponse struct {
	Artists SpotifySearchItems
}

type SpotifySearchItems struct {
	Items []spotify.SpotifyArtist `json:"items"`
	Next string `json:"next"`
}

func (m *musicApi) SpotifySearchArtist(ctx context.Context,
	req *pb.SpotifySearchArtistRequest) (*pb.SpotifySearchArtistResponse, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg(fmt.Sprintf("Start SpotifyFindArtist with req %v", req))

	headers := http.Header{}
	token, ok := ctx.Value(TokenKey).(string)
	if !ok {
		return nil, fmt.Errorf("can't get authorization token")
	}
	headers.Set("Authorization", "Bearer " + token)
	headers.Set("Content-Type", "application/json")

	body := url.Values{}
	query := []string{}
	query = append(query, req.GetName())
	if len(req.GetGenre()) != 0 {
		query = append(query, fmt.Sprintf("genre:%s", strings.Join(req.GetGenre(), ",")))
	}
	body.Add("q", strings.Join(query, " "))
	body.Add("type", "artist")
	body.Add("limit", "1")

	apiRequest, err := http.NewRequest("GET",
		fmt.Sprintf("https://api.spotify.com/v1/search?%s", body.Encode()), nil)

	if err != nil {
		return nil, fmt.Errorf("can't create request to Spotify API, %w", err)
	}
	apiRequest.Header = headers
	log.Debug().Msg(fmt.Sprintf("Successfully create request to Spotify API. req: %v", apiRequest))

	resp, err := http.DefaultClient.Do(apiRequest)
	if err != nil {
		return nil, fmt.Errorf("can't do request to Spotify API, %w", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s can't read response body %w. Status code %d",
			ErrorMessage, err, resp.StatusCode)
	}
	log.Debug().Msg(fmt.Sprintf("Got response from Spotify API. resp: %v and body: %s", resp, data))

	if resp.StatusCode != http.StatusOK {
		jsonResp := &BadSpotifySearchArtist{}
		if err := json.Unmarshal(data, jsonResp); err != nil {
			return nil, fmt.Errorf("%s can't parse response body %w. Status code %d",
				ErrorMessage, err, resp.StatusCode)
		}

		return nil, fmt.Errorf("%s %s. Status code %d", ErrorMessage, jsonResp.Error.Message, resp.StatusCode)
	}

	jsonResp := &OkSpotifySearchResponse{}
	if err := json.Unmarshal(data, jsonResp); err != nil {
		return nil, fmt.Errorf("%s can't parse response body %w. Status code %d",
			ErrorMessage, err, resp.StatusCode)
	}

	if len(jsonResp.Artists.Items) == 0 {
		return nil, custom_errors.ErrArtistNotFound
	}

	foundedArtist := jsonResp.Artists.Items[0]

	return makeSpotifySearchArtistResponse(
		foundedArtist.Name,
		foundedArtist.Followers.Total,
		foundedArtist.Genres,
		foundedArtist.Uri), nil
}

func makeSpotifySearchArtistResponse(
	name string,
	followers int64,
	genres []string,
	uri string) *pb.SpotifySearchArtistResponse {
	return &pb.SpotifySearchArtistResponse {
		Artist: &pb.SpotifyArtist {
			Name: name,
			Followers: followers,
			Genres: genres,
			Uri: uri,
		},
	}
}

package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hihoak/api-service/internal/clients"
	"hihoak/api-service/internal/pkg/utils/logger"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const redirectUriAfterAuthorize = "http://localhost:8080/hello"

type Spotify struct {
	clientId string
	clientSecret string
	redirectUri string
}

func SpotifyFromConfig(cfg *clients.Config) *Spotify {
	return &Spotify{
		clientId: cfg.Spotify.ClientId,
		clientSecret: cfg.Spotify.ClientSecret,
		redirectUri: cfg.Spotify.RedirectUri,
	}
}

type ApiTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
	ExpiresIn uint64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}


func (s Spotify) AuthorizeSpotify(ctx context.Context, w http.ResponseWriter) string {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Start AuthorizeSpotify")

	endpoint := "https://accounts.spotify.com/authorize?"

	data := url.Values{}
	data.Set("client_id", s.clientId)
	data.Set("response_type", "code")
	data.Set("redirect_uri", redirectUriAfterAuthorize)
	state := randStringRunes(16)
	data.Set("state", state)
	data.Set("scope", "user-read-private user-read-email")

	log.Debug().Msg(fmt.Sprintf("Redirecting to Endpoint: '%s', with Data: '%v'", endpoint, data))

	str := endpoint+data.Encode()
	fmt.Println(str)
	http.Redirect(w, &http.Request{}, endpoint+data.Encode(), http.StatusSeeOther)
	return state
}

func (s Spotify) GetSpotifyBearerToken(ctx context.Context, req http.Request, expectedState string) (*ApiTokenResp, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg(fmt.Sprintf("Start GetSpotifyBearerToken. With request: '%v' and expectedState: '%v'",
		req, expectedState))
	parsedURI, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		return nil, err
	}

	errorField, ok := parsedURI.Query()["error"]
	if ok {
		if len(errorField) != 0 {
			return nil, fmt.Errorf("got error while trying to authorize client. Error: '%s'", errorField[0])
		}
		return nil, fmt.Errorf("got error while trying to authorize client")
	}

	currentState, ok := parsedURI.Query()["state"]
	if ! ok {
		return nil, fmt.Errorf("request doesn't contains 'state' field")
	}
	if len(currentState) == 0 {
		return nil, fmt.Errorf("'state' field in request is empty")
	}
	if currentState[0] != expectedState {
		return nil, fmt.Errorf("'state' field in request doesn't match to ")
	}

	codeField, ok := parsedURI.Query()["code"]
	if ! ok {
		return nil, fmt.Errorf("request doesn't contains 'code' field")
	}
	if len(codeField) == 0 {
		return nil, fmt.Errorf("'code' field in request is empty")
	}

	log.Debug().Msg("Validation complete. Start forming POST request")
	data := url.Values{}
	data.Add("grant_type", "authorization_code")
	data.Add("code", codeField[0])
	data.Add("redirect_uri", redirectUriAfterAuthorize)
	body := strings.NewReader(data.Encode())

	postReq, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
	if err != nil {
		return nil, err
	}
	secrets := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", s.clientId, s.clientSecret)))
	postReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", secrets))
	postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	log.Debug().Msg(fmt.Sprintf("Sending POST request to 'https://accounts.spotify.com/api/token' with data: '%v' and headers: '%v'", data, req.Header))
	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Debug().Msg(fmt.Sprintf("Got response %v", resp))

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Debug().Err(err).Msg("Can't read body of response")
		return nil, err
	}

	jsonRespData := ApiTokenResp{}
	if err = json.Unmarshal(respData, &jsonRespData); err != nil {
		log.Debug().Err(err).Msg("Can't parse response body")
	}

	return &jsonRespData, nil
}

package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"hihoak/api-service/internal/clients"
	"hihoak/api-service/internal/clients/spotify"
	"hihoak/api-service/internal/pkg/utils"
	"hihoak/api-service/internal/pkg/utils/logger"
	"hihoak/api-service/internal/server"
	"net/http"
)

func main() {
	ctx := context.Background()

	log.Info().Msg("Starting service...")
	log.Info().Msg("Initializing utils config...")
	utilsCfg, err := utils.LoadConfig(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Some problems with initializing utils config")
	}

	log.Info().Msg("Initializing server config")
	serverCfg, err := server.LoadConfig(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Some problems with initializing server config")
	}

	log.Info().Msg("Initializing clients config")
	clientsConfig, err := clients.LoadConfig(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Some problems with initializing server config")
	}
	spotifyClient := spotify.SpotifyFromConfig(clientsConfig)

	log.Info().Msg("Creating logger")
	globLogger, ctx := logger.FromConfigWithContext(ctx, *utilsCfg)
	globLogger.Info().Msg("Successfully created")

	accessToken := ""
	state := ""

	http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		state = spotifyClient.AuthorizeSpotify(ctx, w)
		globLogger.Debug().Msg(fmt.Sprintf("State of connection: '%s'", state))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		globLogger.Debug().Msg(fmt.Sprintf("Got request: %v", r))
		resp, err := spotifyClient.GetSpotifyBearerToken(ctx, *r, state)
		if err != nil {
			globLogger.Error().Err(err).Msg("Failed to request bearer token")
			return
		}
		accessToken = resp.AccessToken
	})

	go func() {
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			globLogger.Fatal().Err(err).Msg("Can't start http server")
		}
	}()

	grpcServer := server.NewGrpcServer(serverCfg)
	if err := grpcServer.Start(ctx); err != nil {
		globLogger.Fatal().Err(err).Msg("Can't start grpc server")
	}
}

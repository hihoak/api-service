package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"hihoak/api-service/internal/pkg/utils"
	"hihoak/api-service/internal/pkg/utils/logger"
)

func main() {
	ctx := context.Background()

	log.Info().Msg("Starting service")
	log.Info().Msg("Initializing utils config")
	utilsCfg, err := utils.LoadConfig(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Some problems with initializing utils config")
		return
	}

	log.Info().Msg("Creating logger")
	globLogger, ctx := logger.FromConfigWithContext(ctx, *utilsCfg)
	globLogger.Info().Msg("Successfully created")

	return
}

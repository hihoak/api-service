package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hihoak/api-service/internal/api"
	"hihoak/api-service/internal/pkg/utils/logger"
	pb "hihoak/api-service/pkg/pb/hihoak/api-service"
	"net"
)

type grpcServer struct {
	cfg *Config
}

func NewGrpcServer(cfg *Config) *grpcServer {
	return &grpcServer{cfg: cfg}
}

func (g *grpcServer) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg(fmt.Sprintf("Start listening TCP connection on %s:%s", g.cfg.Server.Host, g.cfg.Server.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", g.cfg.Server.Host, g.cfg.Server.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer lis.Close()

	var opts []grpc.ServerOption
	// TODO: read about grpc server options and use something useful
	opts = append(opts, authorizeServerUnaryInterceptor())
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMusicApiServer(grpcServer, api.NewFilmApi())

	err = grpcServer.Serve(lis)
	return fmt.Errorf("Can't start serve, %w", err)
}

func authorizeServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(authorizeInterceptor)
}

func authorizeInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	log := logger.FromContext(ctx)

	var headersToken []string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("empty metadata, pls input bearer token for Spotify API")
	}

	if headersToken, ok = md["authorization-token"]; !ok || len(headersToken) == 0 {
		return nil, fmt.Errorf("needs token to request for Spotify or it must not be empty")
	}

	token := headersToken[0]
	log.Debug().Msg(fmt.Sprintf("Successfully got token and its length %d", len(token)))

	ctx = context.WithValue(ctx, api.TokenKey, token)
	h, err := handler(ctx, req)

	log.Debug().Msg(fmt.Sprintf("Method: %s. Error: %w", info.FullMethod, err))

	return h, err
}

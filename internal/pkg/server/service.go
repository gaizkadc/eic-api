/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package server

import (
	"fmt"
	"github.com/nalej/authx-interceptors/pkg/interceptor/apikey"
	interceptorConfig "github.com/nalej/authx-interceptors/pkg/interceptor/config"
	"github.com/nalej/derrors"
	"github.com/nalej/eic-api/internal/pkg/config"
	"github.com/nalej/eic-api/internal/pkg/server/api"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-eic-api-go"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Service struct {
	Configuration config.Config
}

// NewService creates a new service.
func NewService(conf config.Config) *Service {
	return &Service{
		conf,
	}
}

type Clients struct {
	imClient    grpc_inventory_manager_go.EICClient
	authxClient grpc_authx_go.InventoryClient
}

// GetClients creates the required connections with the remote clients.
func (s *Service) GetClients() (*Clients, derrors.Error) {
	imConn, err := grpc.Dial(s.Configuration.InventoryManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with inventory manager")
	}
	aConn, err := grpc.Dial(s.Configuration.AuthxAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with authx")
	}
	imClient := grpc_inventory_manager_go.NewEICClient(imConn)
	aClient := grpc_authx_go.NewInventoryClient(aConn)
	return &Clients{imClient, aClient}, nil
}

// Run the service, launch the REST service handler.
func (s *Service) Run() error {
	cErr := s.Configuration.Validate()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("invalid configuration")
	}
	s.Configuration.Print()

	authConfig, authErr := s.Configuration.LoadAuthConfig()
	if authErr != nil {
		log.Fatal().Str("err", authErr.DebugReport()).Msg("cannot load authx config")
	}

	clients, cErr := s.GetClients()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("Cannot create clients")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Configuration.Port))
	if err != nil {
		log.Fatal().Errs("failed to listen: %v", []error{err})
	}

	// Create the API Key access provider
	apiKeyAccess := NewAuthxAPIAccess(clients.authxClient)

	// Create handlers
	manager := api.NewManager(clients.imClient)
	handler := api.NewHandler(manager)
	grpcServer := grpc.NewServer(apikey.WithAPIKeyInterceptor(apiKeyAccess,
		interceptorConfig.NewConfig(authConfig, "not-used", s.Configuration.AuthHeader)))
	grpc_eic_api_go.RegisterEICServer(grpcServer, handler)

	if s.Configuration.Debug {
		log.Info().Msg("Enabling gRPC server reflection")
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)
	}
	log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}
	return nil
}

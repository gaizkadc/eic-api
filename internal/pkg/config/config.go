/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package config

import (
	interceptorConfig "github.com/nalej/authx-interceptors/pkg/interceptor/config"
	"github.com/nalej/derrors"
	"github.com/nalej/device-api/version"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// Debug level is active.
	Debug bool
	// Port where the gRPC API service will listen requests.
	Port int
	// InventoryManagerAddress with the host:port to connect to the Inventory manager.
	InventoryManagerAddress string
	// AuthHeader contains the name of the target header.
	AuthHeader string
	// AuthConfigPath contains the path of the file with the authentication configuration.
	AuthConfigPath string
	// AuthxAddress with the host:port to connect to the Authx manager.
	AuthxAddress string
}


func (conf *Config) Validate() derrors.Error {

	if conf.Port <= 0 {
		return derrors.NewInvalidArgumentError("port must be valid")
	}

	if conf.InventoryManagerAddress == "" {
		return derrors.NewInvalidArgumentError("inventoryManager must be set")
	}

	if conf.AuthHeader == "" {
		return derrors.NewInvalidArgumentError("Authorization header must be set")
	}

	if conf.AuthConfigPath == "" {
		return derrors.NewInvalidArgumentError("authConfigPath must be set")
	}

	if conf.AuthxAddress == "" {
		return derrors.NewInvalidArgumentError("authxAddress must be set")
	}

	return nil
}

// LoadAuthConfig loads the security configuration.
func (conf *Config) LoadAuthConfig() (*interceptorConfig.AuthorizationConfig, derrors.Error) {
	return interceptorConfig.LoadAuthorizationConfig(conf.AuthConfigPath)
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.Port).Msg("gRPC port")
	log.Info().Str("URL", conf.InventoryManagerAddress).Msg("Inventory Manager")
	log.Info().Str("URL", conf.AuthxAddress).Msg("Authx")
	log.Info().Str("header", conf.AuthHeader).Msg("Authorization")
	log.Info().Str("path", conf.AuthConfigPath).Msg("Permissions file")
}

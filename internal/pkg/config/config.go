/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	interceptorConfig "github.com/nalej/authx-interceptors/pkg/interceptor/config"
	"github.com/nalej/derrors"
	"github.com/nalej/eic-api/version"
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

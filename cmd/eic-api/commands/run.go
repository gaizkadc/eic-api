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

package commands

import (
	"github.com/nalej/eic-api/internal/pkg/config"
	"github.com/nalej/eic-api/internal/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfg = config.Config{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the Edge Controller API",
	Long:  `Launch the Edge Controller API`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching API!")
		cfg.Debug = debugLevel
		server := server.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&cfg.Port, "port", 5500, "Port to receive management communications")
	runCmd.PersistentFlags().StringVar(&cfg.AuthHeader, "authHeader", "", "Authorization Header")
	runCmd.PersistentFlags().StringVar(&cfg.AuthConfigPath, "authConfigPath", "", "Authorization config path")
	runCmd.PersistentFlags().StringVar(&cfg.InventoryManagerAddress, "inventoryManagerAddress", "", "localhost:6010")
	runCmd.PersistentFlags().StringVar(&cfg.AuthxAddress, "authxAddress", "localhost:8810",
		"Authx address (host:port)")

}

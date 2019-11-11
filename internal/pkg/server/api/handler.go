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

package api

import (
	"context"
	"github.com/nalej/eic-api/internal/pkg/entities"
	"github.com/nalej/grpc-inventory-manager-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	manager Manager
}

func NewHandler(manager Manager) *Handler {
	return &Handler{
		manager: manager,
	}
}

func (h *Handler) Join(ctx context.Context, request *grpc_inventory_manager_go.EICJoinRequest) (*grpc_inventory_manager_go.EICJoinResponse, error) {
	log.Debug().Str("organizationID", request.OrganizationId).Msg("join request")
	verr := entities.ValidEICJoinRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	response, err := h.manager.Join(request)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("organization_id", response.OrganizationId).Str("edge_controller_id", response.EdgeControllerId).Msg("EIC has joined")
	return response, nil
}

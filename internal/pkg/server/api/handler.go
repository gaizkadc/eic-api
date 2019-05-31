/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
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

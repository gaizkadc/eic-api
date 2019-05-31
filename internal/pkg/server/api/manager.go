/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package api

import (
	"context"
	"github.com/nalej/grpc-inventory-manager-go"
	"time"
)

const DefaultInventoryManagerTimeout = time.Second * 30

type Manager struct {
	inventoryManager grpc_inventory_manager_go.EICClient
}

func NewManager(inventoryManager grpc_inventory_manager_go.EICClient) Manager {
	return Manager{
		inventoryManager: inventoryManager,
	}
}

func (m *Manager) Join(request *grpc_inventory_manager_go.EICJoinRequest) (*grpc_inventory_manager_go.EICJoinResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultInventoryManagerTimeout)
	defer cancel()
	return m.inventoryManager.EICJoin(ctx, request)
}

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

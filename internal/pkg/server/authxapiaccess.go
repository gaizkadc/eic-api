/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package server

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-authx-go"
)

type AuthxAPIAccess struct {
	authxClient grpc_authx_go.InventoryClient
}

func NewAuthxAPIAccess(authxClient grpc_authx_go.InventoryClient) * AuthxAPIAccess{
	return &AuthxAPIAccess{
		authxClient:authxClient,
	}
}

func (aa *AuthxAPIAccess) Connect() derrors.Error {
	return nil
}

func (aa *AuthxAPIAccess) IsValid(apiKey string) derrors.Error {
	panic("implement me")
}
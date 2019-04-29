/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package server

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
	"time"
)

const AuthxTimeout = time.Second * 30

// AuthxAPIAccess struct implementing API Key lookup for EIC using the Authx component.
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
	token := &grpc_authx_go.EICJoinToken{
		Token: apiKey,
	}
	ctx, cancel := context.WithTimeout(context.Background(), AuthxTimeout)
	defer cancel()
	_, err := aa.authxClient.ValidEICJoinToken(ctx, token)
	if err != nil{
		derr := conversions.ToDerror(err)
		log.Warn().Str("trace", derr.DebugReport()).Msg("cannot validate token")
		return derrors.NewUnauthenticatedError("cannot validate token")
	}
	return nil
}
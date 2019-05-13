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
	"strings"
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

// token has two field separated by '#'
// the first one is the token and the second one is the organization_id
// we need both to validate the token
func (aa *AuthxAPIAccess) IsValid(tokenInfo string) derrors.Error {

	splitToken := strings.Split(tokenInfo, "#")
	if len(splitToken) != 2{
		log.Warn().Str("tokenInfo", tokenInfo).Msg("cannot validate token. Error in token format")
		return derrors.NewUnauthenticatedError("cannot validate token")
	}
	token := &grpc_authx_go.EICJoinRequest{
		Token: splitToken[0],
		OrganizationId: splitToken[1],
	}
	log.Debug().Interface("token", token).Msg("IsValid")
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
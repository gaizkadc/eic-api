/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-inventory-manager-go"
)

func ValidEICJoinRequest(request *grpc_inventory_manager_go.EICJoinRequest) derrors.Error{
	if request.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id must not be empty")
	}
	return nil
}
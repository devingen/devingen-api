package service_controller

import (
	"context"
	"github.com/devingen/api-core/dvnruntime"
	"github.com/devingen/devingen-api/dto"
	"github.com/devingen/kimlik-api/kimlikruntime"
	"net/http"
)

func (controller ServiceController) GetWorkspaces(ctx context.Context, req dvnruntime.Request) (interface{}, int, error) {

	base, token, err := kimlikruntime.AssertAuthentication(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	workspaceOwnerships, err := controller.Service.WorkspaceOwnershipsFindOfUser(base, token.UserId)
	if err != nil {
		return nil, 0, err
	}

	return &dto.ListResponse{Results: workspaceOwnerships}, http.StatusOK, err
}

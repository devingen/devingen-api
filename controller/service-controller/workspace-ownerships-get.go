package service_controller

import (
	"context"
	core_dto "github.com/devingen/api-core/dto"
	"github.com/devingen/api-core/dvnruntime"
	"github.com/devingen/kimlik-api/kimlikruntime"
	"net/http"
)

func (controller ServiceController) GetWorkspaceOwnerships(ctx context.Context, req dvnruntime.Request) (interface{}, int, error) {

	base, token, err := kimlikruntime.AssertAuthentication(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	workspaceOwnerships, err := controller.Service.FindWorkspaceOwnershipsOfUser(base, token.UserId)
	if err != nil {
		return nil, 0, err
	}

	return &core_dto.GetListResponse{Results: workspaceOwnerships}, http.StatusOK, err
}

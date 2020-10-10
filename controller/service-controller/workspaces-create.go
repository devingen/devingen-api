package service_controller

import (
	"context"
	"github.com/devingen/api-core/dvnruntime"
	"github.com/devingen/api-core/model"
	"github.com/devingen/devingen-api/dto"
	"github.com/devingen/kimlik-api/kimlikruntime"
	"net/http"
)

func (controller ServiceController) CreateWorkspace(ctx context.Context, req dvnruntime.Request) (interface{}, int, error) {

	var body dto.CreateWorkspaceRequest
	base, token, err := kimlikruntime.AssertAuthenticationAndBody(ctx, req, &body)
	if err != nil {
		return nil, 0, err
	}

	if body.Name == "" {
		return nil, 0, model.NewError(http.StatusBadRequest, "name-missing")
	}

	user, err := controller.KimlikService.FindUserUserWithId(base, token.UserId)
	if err != nil {
		return nil, 0, err
	}

	workspace, err := controller.Service.WorkspacesCreate(base, body.Name)
	if err != nil {
		return nil, 0, err
	}

	workspaceOwnership, err := controller.Service.WorkspaceOwnershipsCreate(base, user, workspace)
	if err != nil {
		return nil, 0, err
	}

	return &workspaceOwnership, http.StatusOK, err
}

package service_controller

import (
	"context"
	"github.com/devingen/api-core/dvnruntime"
	coremodel "github.com/devingen/api-core/model"
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

	user, err := controller.KimlikService.FindUserUserWithId(base, token.UserId)
	if err != nil {
		return nil, 0, err
	}

	if user == nil {
		return nil, 0, coremodel.NewError(http.StatusInternalServerError, "user-not-found")
	}

	workspace, err := controller.Service.CreateWorkspace(base, body.Name, user)
	if err != nil {
		return nil, 0, err
	}

	_, err = controller.Service.CreateWorkspaceOwnership(base, user, workspace)
	if err != nil {
		return nil, 0, err
	}

	return &workspace, http.StatusOK, err
}

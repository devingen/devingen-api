package service_controller

import (
	"context"
	"github.com/devingen/api-core/dvnruntime"
	coremodel "github.com/devingen/api-core/model"
	"github.com/devingen/kimlik-api/kimlikruntime"
	"net/http"
)

func (controller ServiceController) GetProductWithId(ctx context.Context, req dvnruntime.Request) (interface{}, int, error) {

	base, token, err := kimlikruntime.AssertAuthentication(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	productId := req.PathParameters["id"]

	product, err := controller.Service.FindProductWithId(base, productId)
	if err != nil {
		return nil, 0, err
	}

	if product == nil {
		return nil, 0, coremodel.NewStatusError(http.StatusNotFound)
	}

	workspaceOwnerships, err := controller.Service.FindWorkspaceOwnershipsOfUser(base, token.UserId)
	if err != nil {
		return nil, 0, err
	}

	found := false
	for _, wo := range workspaceOwnerships {
		if wo.Workspace.ID == product.Workspace.ID {
			found = true
			break
		}
	}

	if !found {
		return nil, 0, coremodel.NewStatusError(http.StatusNotFound)
	}

	return product, http.StatusOK, err
}

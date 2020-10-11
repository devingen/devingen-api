package service_controller

import (
	"context"
	core_dto "github.com/devingen/api-core/dto"
	"github.com/devingen/api-core/dvnruntime"
	"github.com/devingen/kimlik-api/kimlikruntime"
	"net/http"
)

func (controller ServiceController) GetProducts(ctx context.Context, req dvnruntime.Request) (interface{}, int, error) {

	base, token, err := kimlikruntime.AssertAuthentication(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	products, err := controller.Service.FindProductsOfUser(base, token.UserId)
	if err != nil {
		return nil, 0, err
	}

	return &core_dto.GetListResponse{Results: products}, http.StatusOK, err
}

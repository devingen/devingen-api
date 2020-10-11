package dto

import "github.com/devingen/devingen-api/model"

type CreateProductRequest struct {
	Name        string            `json:"name" validate:"min=2,max=32"`
	WorkspaceId string            `json:"workspaceId" validate:"len=24"`
	Type        model.ProductType `json:"type" validate:"oneof=damga"`
}

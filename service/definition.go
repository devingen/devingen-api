package service

import (
	"github.com/devingen/devingen-api/model"
	kimlik_model "github.com/devingen/kimlik-api/model"
)

type IDevingenService interface {
	CreateProduct(base, name string, productType model.ProductType, workspace *model.Workspace, createdBy *kimlik_model.User) (*model.Product, error)
	CreateWorkspace(base, name string, createdBy *kimlik_model.User) (*model.Workspace, error)
	CreateWorkspaceOwnership(base string, user *kimlik_model.User, workspace *model.Workspace) (*model.WorkspaceOwnership, error)
	FindProductsOfUser(base, userId string) ([]*model.Product, error)
	FindProductWithId(base, id string) (*model.Product, error)
	FindWorkspaceOwnershipsOfUser(base, userId string) ([]*model.WorkspaceOwnership, error)
	FindWorkspaceWithId(base, id string) (*model.Workspace, error)
}

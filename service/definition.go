package service

import (
	"github.com/devingen/devingen-api/model"
	kimlik_model "github.com/devingen/kimlik-api/model"
)

type IDevingenService interface {
	WorkspaceOwnershipsFindOfUser(base, userId string) ([]*model.WorkspaceOwnership, error)
	WorkspacesCreate(base, name string) (*model.Workspace, error)
	WorkspaceOwnershipsCreate(base string, user *kimlik_model.User, workspace *model.Workspace) (*model.WorkspaceOwnership, error)
}

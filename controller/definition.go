package controller

import (
	"context"
	"github.com/devingen/api-core/dvnruntime"
)

type IServiceController interface {
	GetWorkspaceOwnerships(ctx context.Context, req dvnruntime.Request) (interface{}, int, error)
	CreateWorkspace(ctx context.Context, req dvnruntime.Request) (interface{}, int, error)
}

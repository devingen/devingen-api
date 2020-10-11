package dto

type CreateWorkspaceRequest struct {
	Name string `json:"name" validate:"min=2,max=32"`
}

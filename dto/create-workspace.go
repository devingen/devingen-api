package dto

type CreateWorkspaceRequest struct {
	Name string `json:"name" validate:"min=6,max=32"`
}

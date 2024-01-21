package models

import "time"

type Workspace struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	OwnerID   int64     `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkspaceCreateReq struct {
	Title   string `json:"title"`
	OwnerID int64  `json:"owner_id"`
}

type WorkspaceUpdateReq struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	OwnerID int64  `json:"owner_id"`
}

type WorkspaceGetReq struct {
	Id int `json:"id"`
}

type WorkspaceFindReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type WorkspaceDeleteReq struct {
	Id int `json:"id"`
}

type WorkspaceFindResponse struct {
	Workspaces []*WorkspaceResponse `json:"workspaces"`
	Count      int                  `json:"count"`
}

type WorkspaceResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	OwnerID   int64     `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkspaceApiResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *WorkspaceResponse
}

type WorkspaceApiFindResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *WorkspaceFindResponse
}

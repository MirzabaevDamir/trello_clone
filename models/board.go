package models

import "time"

type Board struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	OrderNumber int64     `json:"order_number"`
	WorkspaceID int64     `json:"work_space_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BoardCreateReq struct {
	Title       string `json:"title"`
	OrderNumber int64  `json:"order_number"`
	WorkspaceID int64  `json:"work_space_id"`
}

type BoardUpdateReq struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	OrderNumber int64  `json:"order_number"`
	WorkspaceID int64  `json:"work_space_id"`
}

type BoardGetReq struct {
	Id int `json:"id"`
}

type BoardFindReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type BoardDeleteReq struct {
	Id int `json:"id"`
}

type BoardFindResponse struct {
	Boards []*BoardResponse `json:"boards"`
	Count  int              `json:"count"`
}

type BoardResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	OrderNumber int64     `json:"order_number"`
	WorkspaceID int64     `json:"work_space_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BoardApiResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *BoardResponse
}

type BoardApiFindResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *BoardFindResponse
}

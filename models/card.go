package models

import "time"

type Card struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserIDs     []int64   `json:"user_ids"`
	Deadline    time.Time `json:"deadline"`
	OrderNumber int64     `json:"order_number"`
	BoardID     int64     `json:"board_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CardCreateReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserIDs     []int64   `json:"user_ids"`
	Deadline    time.Time `json:"deadline"`
	OrderNumber int64     `json:"order_number"`
	BoardID     int64     `json:"board_id"`
}

type CardUpdateReq struct {
	ID          int64
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserIDs     []int64   `json:"user_ids"`
	Deadline    time.Time `json:"deadline"`
	OrderNumber int64     `json:"order_number"`
	BoardID     int64     `json:"board_id"`
}

type CardGetReq struct {
	Id int `json:"id"`
}

type CardFindReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type CardDeleteReq struct {
	Id int `json:"id"`
}

type CardFindResponse struct {
	Cards []*CardResponse `json:"cards"`
	Count int             `json:"count"`
}

type CardResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserIDs     []int64   `json:"user_ids"`
	Deadline    time.Time `json:"deadline"`
	OrderNumber int64     `json:"order_number"`
	BoardID     int64     `json:"board_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CardApiResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *CardResponse
}

type CardApiFindResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *CardFindResponse
}

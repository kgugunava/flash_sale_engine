package domain

import (
	"time"
)

type Order struct {
	OrderID string
	UserID string
	ItemName string
	Quantity int
	Time time.Time
}

type ResponseStatus struct {
	Success bool
	Code string
	Message string
}

type ResponseError struct {
	Code string
	Message string
	ErrorDetails []string
}

type CreateOrderRequest struct {
	UserID string
	ItemName string
	Quantity int
	Time time.Time
}

type CreateOrderResponse struct {
	Order Order
	Status ResponseStatus
	Error ResponseError
}

type CancelOrderResponse struct {
	Status ResponseStatus
	Error ResponseError
}

type GetOrdersListResponse struct {
	OrdersList []Order
	Status ResponseStatus
	Error ResponseError
}

type GetUserOrdersListResponse struct {
	OrdersList []Order
	Status ResponseStatus
	Error ResponseError
}

type PayForOrderResponse struct {
	Status ResponseStatus
	Error ResponseError
}


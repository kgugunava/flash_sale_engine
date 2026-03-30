package domain

import (
	"time"
)

type Order struct {
	UserID string
	OrderID string
	ItemName string
	Quantity int
	Time time.Time
}

type ResponseError struct {
	Code string
	Message string
	ErrorDetails []string
}

type ResponseStatus struct {
	Success bool
	Code string
	Message string
}

type CreateOrderResponse struct {
	Order Order
	Status ResponseStatus
	Error ResponseError
}
package orders

import (
	"time"
)

type Order struct {
	OrderID string `json:"order_id"`
	UserID string `json:"user_id"`
	ItemName string `json:"item_name"`
	Quantity int `json:"quantity"`
	Time time.Time `json:"time"`
}
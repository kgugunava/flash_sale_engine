package requests

import (
	"time"
)

type OrdersCreateOrderPostRequest struct {
	UserID string `json:"user_id"`
	ItemName string `json:"item_name"`
	Time time.Time `json:"time"`
}
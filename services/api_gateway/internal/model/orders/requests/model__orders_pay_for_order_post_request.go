package requests

import (
	"github.com/google/uuid"
)

type OrdersPayForOrderPostRequest struct {
	OrderID uuid.UUID `uri:"order_id"`
}
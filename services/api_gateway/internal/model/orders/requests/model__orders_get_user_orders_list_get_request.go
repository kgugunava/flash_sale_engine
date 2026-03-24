package requests

import (
	"github.com/google/uuid"
)

type OrdersGetUserOrdersListGetRequest struct {
	UserID uuid.UUID `uri:"user_id"`
}
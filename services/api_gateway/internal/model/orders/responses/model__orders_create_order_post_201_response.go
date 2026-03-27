package responses

import (
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/model/orders"
)

type OrderScreateOrderPost201Response struct {
	Order orders.Order `json:"order"`
}
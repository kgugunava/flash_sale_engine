package requests

type OrdersCreateOrderPostRequest struct {
	UserID string `json:"user_id"`
	ItemName string `json:"item_name"`
	Quantity int `json:"quantity"`
}
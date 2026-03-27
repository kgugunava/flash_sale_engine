package domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	model_orders_requests "github.com/kgugunava/flash_sale_engine/api_gateway/internal/model/orders/requests"
	model_orders "github.com/kgugunava/flash_sale_engine/api_gateway/internal/model/orders"
	model_orders_responses "github.com/kgugunava/flash_sale_engine/api_gateway/internal/model/orders/responses"
	pb_common "github.com/kgugunava/flash_sale_engine/shared/proto/common"
	pb_order "github.com/kgugunava/flash_sale_engine/shared/proto/order"
)

/////// PROTO TO DOMAIN

func ProtoOrderToDomain(protoOrder *pb_order.Order) *Order {
	if protoOrder == nil {
		return nil
	}
	
	return &Order{
		OrderID: protoOrder.OrderId,
		UserID: protoOrder.UserId,
		ItemName: protoOrder.ItemName,
		Quantity: int(protoOrder.Quantity),
		Time: protoOrder.Time.AsTime(),
	}
}

func ProtoResponseStatusToDomain(protoResponseStatus *pb_common.ResponseStatus) *ResponseStatus {
	if protoResponseStatus == nil {
		return nil
	}

	return &ResponseStatus{
		Success: protoResponseStatus.Success,
		Code: protoResponseStatus.Code,
		Message: protoResponseStatus.Message,
	}
}

func ProtoResponseErrorToDomain(protoResponseError *pb_common.ResponseError) *ResponseError {
	if protoResponseError == nil {
		return nil
	}

	return &ResponseError{
		Code: protoResponseError.Code,
		Message: protoResponseError.Message,
		ErrorDetails: protoResponseError.ErrorDetails,
	}
}

func ProtoCreateOrderResponseToDomain(protoResp *pb_order.CreateOrderResponse) *CreateOrderResponse {
	if protoResp == nil {
		return nil
	}

	return &CreateOrderResponse{
		Order: *ProtoOrderToDomain(protoResp.Order),
		Status: *ProtoResponseStatusToDomain(protoResp.Status),
		Error: *ProtoResponseErrorToDomain(protoResp.Error),
	}
}

											/////// JSON TO DOMAIN

func JsonCreateOrderRequestToDomain(req *model_orders_requests.OrdersCreateOrderPostRequest) *CreateOrderRequest {
	if req == nil {
		return nil
	}

	return &CreateOrderRequest{
		UserID: req.UserID,
		ItemName: req.ItemName,
		Quantity: req.Quantity,
	}
}

											/////// DOMAIN TO PROTO

func DomainCreateOrderRequestToProto(req *CreateOrderRequest) *pb_order.CreateOrderRequest {
	if req == nil {
		return nil
	}

	return &pb_order.CreateOrderRequest{
		UserId: req.UserID,
		ItemName: req.ItemName,
		Quantity: int32(req.Quantity),
		Time: timestamppb.New(req.Time),
	}
}

											/////// DOMAIN TO JSON

func DomainOrderToJson(domainOrder *Order) *model_orders.Order {
	if domainOrder == nil {
		return nil
	}

	return &model_orders.Order{
		OrderID: domainOrder.OrderID,
		UserID: domainOrder.UserID,
		ItemName: domainOrder.ItemName,
		Quantity: domainOrder.Quantity,
		Time: domainOrder.Time,
	}
}

func DomainCreateOrderResponseToJson201Response(domainResp *CreateOrderResponse) *model_orders_responses.OrderScreateOrderPost201Response {
	if domainResp == nil {
		return nil
	}

	return &model_orders_responses.OrderScreateOrderPost201Response{
		Order: *DomainOrderToJson(&domainResp.Order),
	}
}
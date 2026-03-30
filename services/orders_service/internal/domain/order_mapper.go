package domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	pb_order "github.com/kgugunava/flash_sale_engine/shared/proto/order"
	pb_common "github.com/kgugunava/flash_sale_engine/shared/proto/common"
)

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

func ProtoCreateOrderRequestToDomainOrder(req *pb_order.CreateOrderRequest) *Order {
	if req == nil {
		return nil
	}

	return &Order{
		UserID: req.UserId,
		ItemName: req.ItemName,
		Quantity: int(req.Quantity),
		Time: req.Time.AsTime(),
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

func DomainOrderToProto(domainOrder *Order) *pb_order.Order {
	if domainOrder == nil {
		return nil
	}

	return &pb_order.Order{
		OrderId: domainOrder.OrderID,
		UserId: domainOrder.UserID,
		ItemName: domainOrder.ItemName,
		Quantity: int32(domainOrder.Quantity),
		Time: timestamppb.New(domainOrder.Time),
	}
}
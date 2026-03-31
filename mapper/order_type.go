package mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen/v1"
)

func OrderTypeFromProto(orderType pb.OrderType) string {
	switch orderType {
	case pb.OrderType_ORDER_TYPE_BUY:
		return "BUY"
	case pb.OrderType_ORDER_TYPE_SELL:
		return "SELL"
	default:
		return "UNSPECIFIED"
	}
}

func OrderTypeToProto(orderType string) pb.OrderType {
	switch orderType {
	case "BUY":
		return pb.OrderType_ORDER_TYPE_BUY
	case "SELL":
		return pb.OrderType_ORDER_TYPE_SELL
	default:
		return pb.OrderType_ORDER_TYPE_UNSPECIFIED
	}
}

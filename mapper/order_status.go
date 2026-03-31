package mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen/v1"
)

func OrderStatusFromProto(orderStatus pb.OrderStatus) string {
	switch orderStatus {
	case pb.OrderStatus_ORDER_STATUS_CREATED:
		return "CREATED"
	case pb.OrderStatus_ORDER_STATUS_PENDING:
		return "PENDING"
	case pb.OrderStatus_ORDER_STATUS_WAIT_SELLER:
		return "WAIT_SELLER"
	case pb.OrderStatus_ORDER_STATUS_PAID:
		return "PAID"
	case pb.OrderStatus_ORDER_STATUS_ON_HOLD:
		return "ON_HOLD"
	case pb.OrderStatus_ORDER_STATUS_PROCESSING:
		return "PROCESSING"
	case pb.OrderStatus_ORDER_STATUS_PACKED:
		return "PACKED"
	case pb.OrderStatus_ORDER_STATUS_OUT_OF_DELIVERY:
		return "OUT_OF_DELIVERY"
	case pb.OrderStatus_ORDER_STATUS_ON_THE_WAY:
		return "ORDER_STATUS_ON_THE_WAY"
	case pb.OrderStatus_ORDER_STATUS_DELIVERED:
		return "DELIVERED"
	case pb.OrderStatus_ORDER_STATUS_CLOSED:
		return "CLOSED"
	default:
		return "UNSPECIFIED"
	}
}

func OrderStatusToProto(orderStatus string) pb.OrderStatus {
	switch orderStatus {
	case "CREATED":
		return pb.OrderStatus_ORDER_STATUS_CREATED
	case "PENDING":
		return pb.OrderStatus_ORDER_STATUS_PENDING
	case "WAIT_SELLER":
		return pb.OrderStatus_ORDER_STATUS_WAIT_SELLER
	case "PAID":
		return pb.OrderStatus_ORDER_STATUS_PAID
	case "ON_HOLD":
		return pb.OrderStatus_ORDER_STATUS_ON_HOLD
	case "PROCESSING":
		return pb.OrderStatus_ORDER_STATUS_PROCESSING
	case "PACKED":
		return pb.OrderStatus_ORDER_STATUS_PACKED
	case "OUT_OF_DELIVERY":
		return pb.OrderStatus_ORDER_STATUS_OUT_OF_DELIVERY
	case "ON_THE_WAY":
		return pb.OrderStatus_ORDER_STATUS_ON_THE_WAY
	case "DELIVERED":
		return pb.OrderStatus_ORDER_STATUS_DELIVERED
	case "CLOSED":
		return pb.OrderStatus_ORDER_STATUS_CLOSED
	default:
		return pb.OrderStatus_ORDER_STATUS_UNSPECIFIED
	}
}

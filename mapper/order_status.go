package mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen/v2"
)

func OrderStatusFromProto(orderStatus pb.OrderStatus) string {
	switch orderStatus {
	case 1:
		return "CREATED"
	case 2:
		return "PENDING"
	case 3:
		return "WAIT_SELLER"
	case 4:
		return "PAID"
	case 5:
		return "ON_HOLD"
	case 6:
		return "PROCESSING"
	case 7:
		return "PACKED"
	case 8:
		return "OUT_OF_DELIVERY"
	case 9:
		return "ORDER_STATUSON_THE_WAY"
	case 10:
		return "DELIVERED"
	case 11:
		return "CLOSED"
	default:
		return "UNSPECIFIED"
	}
}

func OrderStatusToProto(orderStatus string) pb.OrderStatus {
	switch orderStatus {
	case "CREATED":
		return 1
	case "PENDING":
		return 2
	case "WAIT_SELLER":
		return 3
	case "PAID":
		return 4
	case "ON_HOLD":
		return 5
	case "PROCESSING":
		return 6
	case "PACKED":
		return 7
	case "OUT_OF_DELIVERY":
		return 8
	case "ORDER_STATUSON_THE_WAY":
		return 9
	case "DELIVERED":
		return 10
	case "CLOSED":
		return 11
	default:
		return 0
	}
}

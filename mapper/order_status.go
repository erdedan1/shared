package proto_mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen"
)

func StringOrderStatusToProto(strOrderStatus string) pb.OrderStatus {
	if r, ok := pb.OrderStatus_value[strOrderStatus]; ok {
		return pb.OrderStatus(r)
	}
	return pb.OrderStatus_ORDER_STATUS_UNSPECIFIED
}

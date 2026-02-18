package proto_mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen"
)

func ProtoOrderStatusToString(protoOrderStatus []pb.OrderStatus) []string {
	orderStatus := make([]string, 0, len(protoOrderStatus))
	for _, status := range protoOrderStatus {
		orderStatus = append(orderStatus, status.String())
	}
	return orderStatus
}

func StringOrderStatusToProto(strOrderStatus []string) []pb.OrderStatus {
	orderStatus := make([]pb.OrderStatus, 0, len(strOrderStatus))
	for _, status := range strOrderStatus {
		if r, ok := pb.OrderStatus_value[status]; ok {
			orderStatus = append(orderStatus, pb.OrderStatus(r))
		}
	}
	return orderStatus
}

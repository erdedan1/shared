package proto_mapper

import (
	"errors"

	pb "github.com/erdedan1/protocol/proto/order_service/gen"
)

func StringOrderStatusToProto(strOrderStatus string) (*pb.OrderStatus, error) {
	if r, ok := pb.OrderStatus_value[strOrderStatus]; ok {
		return pb.OrderStatus(r).Enum(), nil
	}
	return nil, errors.New("")
}

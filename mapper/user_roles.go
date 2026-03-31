package mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen/v1"
)

func UserRoleFromProto(userRole pb.UserRole) string {
	switch userRole {
	case pb.UserRole_USER_ROLE_TRADER:
		return "TRADER"
	case pb.UserRole_USER_ROLE_ADMIN:
		return "ADMIN"
	case pb.UserRole_USER_ROLE_VIEWER:
		return "VIEWER"
	default:
		return "UNSPECIFIED"
	}
}

func UserRoleToProto(userRole string) pb.UserRole {
	switch userRole {
	case "TRADER":
		return pb.UserRole_USER_ROLE_TRADER
	case "ADMIN":
		return pb.UserRole_USER_ROLE_ADMIN
	case "VIEWER":
		return pb.UserRole_USER_ROLE_VIEWER
	default:
		return pb.UserRole_USER_ROLE_UNSPECIFIED
	}
}

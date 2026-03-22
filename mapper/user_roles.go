package mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen"
)

func UserRoleFromProto(userRole pb.UserRole) string {
	switch userRole {
	case 1:
		return "TRADER"
	case 2:
		return "ADMIN"
	case 3:
		return "VIEWER"
	default:
		return "UNSPECIFIED"
	}
}

func UserRoleToProto(userRole string) pb.UserRole {
	switch userRole {
	case "TRADER":
		return 1
	case "ADMIN":
		return 2
	case "VIEWER":
		return 3
	default:
		return 0
	}
}

func UserRolesToProto(urs []string) []pb.UserRole {
	res := make([]pb.UserRole, 0, len(urs))
	for _, role := range urs {
		pbRole := UserRoleToProto(role)
		res = append(res, pbRole)
	}
	return res
}

func UserRolesFromProto(pbRoles []pb.UserRole) []string {
	res := make([]string, 0, len(pbRoles))
	for _, role := range pbRoles {
		urs := UserRoleFromProto(role)
		res = append(res, urs)
	}
	return res
}

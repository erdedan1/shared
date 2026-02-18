package proto_mapper

import (
	pb "github.com/erdedan1/protocol/proto/order_service/gen"
)

func ProtoUserRolesToString(protoUserRoles []pb.UserRole) []string {
	userRoles := make([]string, 0, len(protoUserRoles))
	for _, role := range protoUserRoles {
		userRoles = append(userRoles, role.String())
	}
	return userRoles
}

func StringUserRolesToProto(strUserRoles []string) []pb.UserRole {
	userRoles := make([]pb.UserRole, 0, len(strUserRoles))
	for _, role := range strUserRoles {
		if r, ok := pb.UserRole_value[role]; ok {
			userRoles = append(userRoles, pb.UserRole(r))
		}
	}
	return userRoles
}

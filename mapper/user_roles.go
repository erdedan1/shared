package mapper

import (
	pbo "github.com/erdedan1/protocol/proto/order_service/gen/v1"
	pbs "github.com/erdedan1/protocol/proto/spot_instrument_service/gen/v1"
)

func UserRoleFromProtoSpot(userRole pbs.UserRole) string {
	switch userRole {
	case pbs.UserRole_USER_ROLE_TRADER:
		return "TRADER"
	case pbs.UserRole_USER_ROLE_ADMIN:
		return "ADMIN"
	case pbs.UserRole_USER_ROLE_VIEWER:
		return "VIEWER"
	default:
		return "UNSPECIFIED"
	}
}

func UserRoleToProtoSpot(userRole string) pbs.UserRole {
	switch userRole {
	case "TRADER":
		return pbs.UserRole_USER_ROLE_TRADER
	case "ADMIN":
		return pbs.UserRole_USER_ROLE_ADMIN
	case "VIEWER":
		return pbs.UserRole_USER_ROLE_VIEWER
	default:
		return pbs.UserRole_USER_ROLE_UNSPECIFIED
	}
}

func UserRoleFromProtoOrder(userRole pbo.UserRole) string {
	switch userRole {
	case pbo.UserRole_USER_ROLE_TRADER:
		return "TRADER"
	case pbo.UserRole_USER_ROLE_ADMIN:
		return "ADMIN"
	case pbo.UserRole_USER_ROLE_VIEWER:
		return "VIEWER"
	default:
		return "UNSPECIFIED"
	}
}

func UserRoleToProtoOrder(userRole string) pbo.UserRole {
	switch userRole {
	case "TRADER":
		return pbo.UserRole_USER_ROLE_TRADER
	case "ADMIN":
		return pbo.UserRole_USER_ROLE_ADMIN
	case "VIEWER":
		return pbo.UserRole_USER_ROLE_VIEWER
	default:
		return pbo.UserRole_USER_ROLE_UNSPECIFIED
	}
}

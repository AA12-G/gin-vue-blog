package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin         Role = 1 //管理员
	PermissionUser          Role = 2 //普通登录人
	PermissionVisitor       Role = 3 //游客
	PermissionSuperDisabled Role = 4 //被禁用的用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	switch s {
	case PermissionAdmin:
		return "管理员"

	case PermissionUser:
		return "普通用户"

	case PermissionVisitor:
		return "游客"

	case PermissionSuperDisabled:
		return "被禁用"

	default:
		return "其他"
	}
}

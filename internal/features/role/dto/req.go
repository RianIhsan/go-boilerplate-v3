package dto

import "ams-sentuh/internal/entities"

type RegisterRoleRequest struct {
	Name        string       `json:"name" validate:"required,min=3,max=50"`
	Permissions []Permission `json:"permissions"`
	Access      []Access     `json:"access"`
	ListAction  []ListAction `json:"list_action"`
}
type Permission struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	AccessId uint64 `json:"access_id"`
}
type Access struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RolePermissions struct {
	RoleId       uint         `json:"role_id"`
	PermissionId uint64       `json:"permission_id,omitempty"`
	AccessId     uint         `json:"access_id,omitempty"`
	ListAction   []ListAction `json:"list_action"`
}
type ListAction struct {
	ID     uint64 `json:"id"`
	Action int8   `json:"action"`
	Method string `json:"method,omitempty"`
	Path   string `json:"path,omitempty"`
}

func ConvertToEntityRoleRequest(req RegisterRoleRequest) entities.Role {
	return entities.Role{
		Name: req.Name,
	}
}

type UpdateRoleRequest struct {
	Name string
}

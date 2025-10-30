package dto

import (
	"ams-sentuh/internal/entities"
	"time"
)

type RegisterRoleResponse struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type RoleResponse struct {
	ID          uint64       `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	Access      []Access     `json:"accesses"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   *time.Time   `json:"deleted_at,omitempty"`
}

func ConvertToGetAllRoleResponse(data entities.Role) RoleResponse {
	return RoleResponse{
		ID:          data.ID,
		Name:        data.Name,
		Permissions: ToListPermission(data.Permissions),
		Access:      ToListAccessCore(data.Accesses),
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeletedAt:   data.DeletedAt,
	}
}

func ConvertToGetListRoleResponse(data []entities.Role) []RoleResponse {
	var result []RoleResponse
	for _, role := range data {
		result = append(result, ConvertToGetAllRoleResponse(role))
	}
	return result
}

func ConvertToRegisterRoleResponse(role entities.Role) RegisterRoleResponse {
	var deletedAt *string
	if role.DeletedAt != nil {
		deletedAtStr := role.DeletedAt.Format("2006-01-02 15:04:05")
		deletedAt = &deletedAtStr
	}

	return RegisterRoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func ConvertToRoleResponse(roles entities.Role) RoleResponse {
	return RoleResponse{
		ID:          roles.ID,
		Name:        roles.Name,
		Permissions: ToListPermission(roles.Permissions),
		Access:      ToListAccessCore(roles.Accesses),
		CreatedAt:   roles.CreatedAt,
		UpdatedAt:   roles.UpdatedAt,
	}
}

func ConvertPermissionResponse(data entities.Permission) Permission {
	return Permission{
		ID:       uint(data.ID),
		Name:     data.Name,
		Path:     data.Path,
		Method:   data.Method,
		AccessId: *data.AccessID,
	}
}

func ConvertAccessResponse(data entities.Access) Access {
	return Access{
		ID:   uint(data.ID),
		Name: data.Name,
	}
}

func ToListPermission(data []entities.Permission) []Permission {
	var result []Permission
	for _, permission := range data {
		result = append(result, ConvertPermissionResponse(permission))
	}
	return result
}

func ToListAccessCore(data []entities.Access) []Access {
	var result []Access
	for _, access := range data {
		result = append(result, ConvertAccessResponse(access))
	}
	return result
}

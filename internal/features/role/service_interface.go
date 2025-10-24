package role

import (
	"context"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/role/dto"
)

type RoleServiceInterface interface {
	AddRole(ctx context.Context, request dto.RegisterRoleRequest) (dto.RegisterRoleResponse, error)
	GetAll(ctx context.Context) ([]dto.RoleResponse, error)
	ModifyRolePermission(ctx context.Context, data dto.RolePermissions) (entities.Role, error)
	GetByID(ctx context.Context, roleId uint64) (dto.RoleResponse, error)
	UpdateRole(ctx context.Context, id uint64, req dto.UpdateRoleRequest) error
	DeleteRole(ctx context.Context, roleId uint64) error
}

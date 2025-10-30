package role

import (
	"context"
	"ams-sentuh/internal/entities"
)

type RoleRepositoryInterface interface {
	Create(ctx context.Context, data entities.Role) (entities.Role, error)
	GetAll(ctx context.Context) ([]entities.Role, error)
	GetByID(ctx context.Context, roleID uint) (*entities.Role, error)
	UpdateRole(ctx context.Context, role *entities.Role) error
	DeleteRole(ctx context.Context, id uint64) error
	SearchRoleByName(ctx context.Context, name string) ([]entities.Role, error)
	UpdateRoleAccess(ctx context.Context, roleID uint64, accessList []entities.Access) error
	CreateRoleAccess(ctx context.Context, accessList []entities.RoleAccess) error
	DeleteRoleAccess(ctx context.Context, roleID uint64, accessIDs []uint64) error
	GetRoleAccess(ctx context.Context, roleID uint64) ([]entities.RoleAccess, error)
	UpdateRolePermission(ctx context.Context, roleID uint64, permissions []entities.Permission) error
	CreateRolePermission(ctx context.Context, perms []entities.RolePermission) error
	DeleteRolePermission(ctx context.Context, roleID uint64, permIDs []uint64) error
	DeletePermissionByAccessNotIn(ctx context.Context, roleID uint64, validAccessIDs []uint64) error
	GetPermissionByAccess(ctx context.Context, accessID uint64) ([]entities.Permission, error)
	BatchGetPermissions(ctx context.Context, permIDs []uint64) ([]entities.Permission, error)
	GetGroupedPermissions(ctx context.Context, roleID uint64) ([]entities.RolePermission, error)
}

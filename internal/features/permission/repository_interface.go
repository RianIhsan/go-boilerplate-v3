package permission

import (
	"context"
	"ams-sentuh/internal/entities"
)

type PermissionRepositoryInterface interface {
	Create(ctx context.Context, data entities.Permission) (entities.Permission, error)
	DeletePermission(ctx context.Context, id uint64) error
	UpdatePermission(ctx context.Context, data entities.Permission) (entities.Permission, error)
	GetPermission(ctx context.Context, id uint64) (entities.Permission, error)
	GetPermissions(ctx context.Context, accessID *uint64) ([]entities.Permission, error)
}

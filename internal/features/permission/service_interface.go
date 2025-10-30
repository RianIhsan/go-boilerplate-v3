package permission

import (
	"context"
	"ams-sentuh/internal/features/permission/dto"
)

type PermissionServiceInterface interface {
	AddPermission(ctx context.Context, register dto.PermissionRegister) (dto.Permission, error)
	GetListPermission(ctx context.Context, accessID *uint64) ([]dto.Permission, error)
	GetPermission(ctx context.Context, id uint64) (dto.Permission, error)
	UpdatePermission(ctx context.Context, id uint64, update dto.PermissionUpdate) error
	DeletePermission(ctx context.Context, id uint64) error
}

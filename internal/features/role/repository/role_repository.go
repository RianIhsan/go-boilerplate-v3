package repository

import (
	"context"
	"gorm.io/gorm"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/role"
	"time"
)

type rolePostgresRepository struct {
	db *gorm.DB
}

func NewRolePostgresRepository(db *gorm.DB) role.RoleRepositoryInterface {
	return &rolePostgresRepository{
		db: db,
	}
}

// ===== Role CRUD =====

func (r *rolePostgresRepository) Create(ctx context.Context, data entities.Role) (entities.Role, error) {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return entities.Role{}, err
	}
	return data, nil
}

func (r *rolePostgresRepository) GetByID(ctx context.Context, id uint) (*entities.Role, error) {
	var dataRole entities.Role
	err := r.db.WithContext(ctx).
		Preload("Permissions").
		Preload("Accesses").
		First(&dataRole, id).Error
	return &dataRole, err
}

func (r *rolePostgresRepository) UpdateRole(ctx context.Context, role *entities.Role) error {
	role.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).
		Model(&entities.Role{}).
		Where("id = ?", role.ID).
		Updates(role).Error
}

func (r *rolePostgresRepository) DeleteRole(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).
		Delete(&entities.Role{}, id).Error
}

func (r *rolePostgresRepository) GetAll(ctx context.Context) ([]entities.Role, error) {
	var roles []entities.Role
	err := r.db.WithContext(ctx).
		Preload("Permissions").
		Preload("Accesses").
		Where("deleted_at IS NULL").
		Find(&roles).Error
	return roles, err
}

func (r *rolePostgresRepository) SearchRoleByName(ctx context.Context, name string) ([]entities.Role, error) {
	var roles []entities.Role
	err := r.db.WithContext(ctx).
		Where("name ILIKE ?", "%"+name+"%").
		Find(&roles).Error
	return roles, err
}

// ===== Access Management =====

func (r *rolePostgresRepository) UpdateRoleAccess(ctx context.Context, roleID uint64, accessList []entities.Access) error {
	role := entities.Role{ID: roleID}
	return r.db.WithContext(ctx).
		Model(&role).
		Association("Access").
		Replace(accessList)
}

func (r *rolePostgresRepository) CreateRoleAccess(ctx context.Context, accessList []entities.RoleAccess) error {
	return r.db.WithContext(ctx).
		Table("role_accesses").
		Create(&accessList).Error
}

func (r *rolePostgresRepository) DeleteRoleAccess(ctx context.Context, roleID uint64, accessIDs []uint64) error {
	return r.db.WithContext(ctx).
		Table("role_accesses").
		Where("role_id = ? AND access_id IN ?", roleID, accessIDs).
		Delete(&entities.RoleAccess{}).Error
}

func (r *rolePostgresRepository) GetRoleAccess(ctx context.Context, roleID uint64) ([]entities.RoleAccess, error) {
	var result []entities.RoleAccess
	err := r.db.WithContext(ctx).
		Table("role_accesses").
		Where("role_id = ?", roleID).
		Find(&result).Error
	return result, err
}

// ===== Permission Management =====

func (r *rolePostgresRepository) UpdateRolePermission(ctx context.Context, roleID uint64, permissions []entities.Permission) error {
	role := entities.Role{ID: roleID}
	return r.db.WithContext(ctx).
		Model(&role).
		Association("Permissions").
		Replace(permissions)
}

func (r *rolePostgresRepository) CreateRolePermission(ctx context.Context, perms []entities.RolePermission) error {
	return r.db.WithContext(ctx).
		Create(&perms).Error
}

func (r *rolePostgresRepository) DeleteRolePermission(ctx context.Context, roleID uint64, permIDs []uint64) error {
	return r.db.WithContext(ctx).
		Table("role_permissions").
		Where("role_id = ? AND permission_id IN ?", roleID, permIDs).
		Delete(&entities.RolePermission{}).Error
}

func (r *rolePostgresRepository) DeletePermissionByAccessNotIn(ctx context.Context, roleID uint64, validAccessIDs []uint64) error {
	return r.db.WithContext(ctx).
		Table("role_permissions").
		Where("role_id = ? AND access_id NOT IN ?", roleID, validAccessIDs).
		Delete(&entities.RolePermission{}).Error
}

func (r *rolePostgresRepository) GetPermissionByAccess(ctx context.Context, accessID uint64) ([]entities.Permission, error) {
	var perms []entities.Permission
	err := r.db.WithContext(ctx).
		Where("access_id = ?", accessID).
		Find(&perms).Error
	return perms, err
}

func (r *rolePostgresRepository) BatchGetPermissions(ctx context.Context, permIDs []uint64) ([]entities.Permission, error) {
	var perms []entities.Permission
	err := r.db.WithContext(ctx).
		Where("id IN ?", permIDs).
		Find(&perms).Error
	return perms, err
}

// ===== Helper: Grouped Permissions =====

func (r *rolePostgresRepository) GetGroupedPermissions(ctx context.Context, roleID uint64) ([]entities.RolePermission, error) {
	var list []entities.RolePermission
	err := r.db.WithContext(ctx).
		Table("role_permissions").
		Select("role_id, access_id").
		Where("role_id = ?", roleID).
		Group("role_id, access_id").
		Find(&list).Error
	return list, err
}

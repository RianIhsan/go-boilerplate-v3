package repository

import (
	"context"
	"fmt"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/permission"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type permissionPostgresRepository struct {
	db *gorm.DB
}

func NewPermissionPostgresRepository(db *gorm.DB) permission.PermissionRepositoryInterface {
	return &permissionPostgresRepository{
		db: db,
	}
}

func (p permissionPostgresRepository) Create(ctx context.Context, data entities.Permission) (entities.Permission, error) {
	if err := p.db.Create(&data).Error; err != nil {
		return entities.Permission{}, err
	}
	return data, nil
}

func (p permissionPostgresRepository) DeletePermission(ctx context.Context, id uint64) error {
	db := p.db.WithContext(ctx)

	var permission entities.Permission
	if err := db.First(&permission, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("permission not found")
		}
		return err
	}
	if err := db.Where("id = ?", id).Delete(&entities.Permission{}).Error; err != nil {
		return err
	}
	return nil
}

func (p permissionPostgresRepository) UpdatePermission(ctx context.Context, data entities.Permission) (entities.Permission, error) {
	db := p.db.WithContext(ctx)

	var existingPermission entities.Permission
	if err := db.First(&existingPermission, data.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Permission{}, errors.New("permission not found")
		}
		return entities.Permission{}, err
	}

	// Validasi AccessID jika tidak nil
	if data.AccessID != nil {
		if *data.AccessID == 0 {
			// Jika AccessID adalah 0, set ke nil
			data.AccessID = nil
		} else {
			// Validasi apakah AccessID exists di tabel accesses
			var accessExists bool
			if err := db.Model(&entities.Access{}).
				Select("count(*) > 0").
				Where("id = ?", *data.AccessID).
				Find(&accessExists).Error; err != nil {
				return entities.Permission{}, fmt.Errorf("failed to check access existence: %w", err)
			}

			if !accessExists {
				return entities.Permission{}, fmt.Errorf("access with id %d does not exist", *data.AccessID)
			}
		}
	}

	// Selective update - hanya update field yang tidak zero value
	updateData := make(map[string]interface{})

	if data.Name != "" {
		updateData["name"] = data.Name
	}
	if data.Path != "" {
		updateData["path"] = data.Path
	}
	if data.Method != "" {
		updateData["method"] = data.Method
	}
	if data.Type != "" {
		updateData["type"] = data.Type
	}

	// Handle AccessID secara khusus
	if data.AccessID != nil {
		if *data.AccessID == 0 {
			updateData["access_id"] = nil
		} else {
			updateData["access_id"] = *data.AccessID
		}
	}

	// Update dengan map untuk menghindari zero value
	if len(updateData) > 0 {
		updateData["updated_at"] = time.Now()
		if err := db.Model(&existingPermission).Updates(updateData).Error; err != nil {
			return entities.Permission{}, fmt.Errorf("failed to update permission: %w", err)
		}
	}

	// Return updated permission dengan data terbaru
	var updatedPermission entities.Permission
	if err := db.First(&updatedPermission, data.ID).Error; err != nil {
		return entities.Permission{}, fmt.Errorf("failed to get updated permission: %w", err)
	}

	return updatedPermission, nil
}

func (p permissionPostgresRepository) GetPermission(ctx context.Context, id uint64) (entities.Permission, error) {
	var permission entities.Permission
	if err := p.db.WithContext(ctx).First(&permission, id).Error; err != nil {
		return entities.Permission{}, err
	}
	return permission, nil
}

func (p permissionPostgresRepository) GetPermissions(ctx context.Context, accessID *uint64) ([]entities.Permission, error) {
	var permissions []entities.Permission
	query := p.db.WithContext(ctx).Model(&entities.Permission{})

	if accessID != nil {
		query = query.Where("access_id = ?", *accessID)
	}

	if err := query.Order("id ASC").Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

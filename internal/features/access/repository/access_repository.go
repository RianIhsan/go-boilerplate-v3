package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/access"
)

type accessPostgresRepository struct {
	db *gorm.DB
}

func NewAccessPostgresRepository(db *gorm.DB) access.AccessRepositoryInterface {
	return &accessPostgresRepository{db: db}
}

func (ac *accessPostgresRepository) Create(ctx context.Context, model entities.Access) (entities.Access, error) {
	db := ac.db.WithContext(ctx)
	if err := db.Create(&model).Error; err != nil {
		return entities.Access{}, err
	}
	return model, nil
}

func (ac *accessPostgresRepository) GetAll(ctx context.Context) ([]entities.Access, error) {
	db := ac.db.WithContext(ctx)
	var accesses []entities.Access

	if err := db.Preload("Permissions").Order("id ASC").Find(&accesses).Error; err != nil {
		return nil, err
	}

	return accesses, nil
}

func (ac *accessPostgresRepository) GetByID(ctx context.Context, id uint64) (entities.Access, error) {
	db := ac.db.WithContext(ctx)
	var access entities.Access
	if err := db.Preload("Permissions").First(&access, id).Error; err != nil {
		return entities.Access{}, err
	}
	return access, nil
}

func (ac *accessPostgresRepository) Update(ctx context.Context, model entities.Access) (entities.Access, error) {
	db := ac.db.WithContext(ctx)

	var existingAccess entities.Access
	if err := db.First(&existingAccess, model.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Access{}, errors.New("access not found")
		}
		return entities.Access{}, err
	}

	if err := db.Model(&existingAccess).Updates(model).Error; err != nil {
		return entities.Access{}, err
	}

	return existingAccess, nil
}

func (ac *accessPostgresRepository) Delete(ctx context.Context, id uint64) error {
	db := ac.db.WithContext(ctx)

	var access entities.Access
	if err := db.First(&access, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("access not found")
		}
		return err
	}

	if err := db.Where("id = ?", id).Delete(&access).Error; err != nil {
		return err
	}

	return nil
}

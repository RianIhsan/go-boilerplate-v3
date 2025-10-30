package db

import (
	"ams-sentuh/internal/entities"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Role{},
		&entities.Access{},
		&entities.Permission{},
		&entities.RoleAccess{},
		&entities.RolePermission{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}

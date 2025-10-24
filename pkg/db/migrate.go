package db

import (
	"fmt"
	"ams-sentuh/internal/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
		// &entities.Device{},
		&entities.Role{},
		&entities.Permission{},
		&entities.Access{},
		&entities.RoleAccess{},
		&entities.RolePermission{},
		// &entities.DeviceAddon{},
		// &entities.DeviceApplication{},
		// &entities.DeviceType{},
		// &entities.DeviceFeature{},
		// &entities.DeviceSpecification{},
		// &entities.DeviceTypeFeature{},
		// &entities.WhitelistApp{},
		// &entities.Log{},
		// &entities.Area{},
		// &entities.Province{},
		// &entities.City{},
		// &entities.Branch{},
		// &entities.AreaPhoto{},
		// &entities.Client{},
		// &entities.Ticket{},
		// &entities.TicketAttachment{},
		// &entities.TicketAssignment{},
		// &entities.OTP{},
		// &entities.Company{},
		// &entities.DeviceHealth{},
		// &entities.Widget{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}

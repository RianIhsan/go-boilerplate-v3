package entities

import "time"

type Role struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Permissions []Permission `gorm:"many2many:role_permissions;joinForeignKey:RoleID;joinReferences:PermissionID" json:"permissions,omitempty"`

	Accesses  []Access   `gorm:"many2many:role_accesses;joinForeignKey:RoleID;joinReferences:AccessID" json:"accesses,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index;"`
}

type RolePermission struct {
	RoleID       uint64 `gorm:"primaryKey"`
	PermissionID uint64 `gorm:"primaryKey"`
	AccessID     *uint64
}

type RoleAccess struct {
	AccessID uint64 `gorm:"primaryKey"`
	RoleID   uint64 `gorm:"primaryKey"`

	Access Access `gorm:"foreignKey:AccessID"`
	Role   Role   `gorm:"foreignKey:RoleID"`
}

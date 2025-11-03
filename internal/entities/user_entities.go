package entities

import "time"

type User struct {
	ID        uint64  `gorm:"primaryKey;autoIncrement"`
	Avatar    string  `gorm:"type:varchar(255)"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Username  string  `gorm:"type:varchar(100);not null"`
	Email     string  `gorm:"type:varchar(100);not null;uniqueIndex:uni_users_email"`
	Password  string  `gorm:"type:varchar(100);not null"`
	NFCTag    *string `gorm:"type:varchar(100)"`
	RoleID    uint64
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`

	Role *Role `gorm:"foreignKey:RoleID"`
}

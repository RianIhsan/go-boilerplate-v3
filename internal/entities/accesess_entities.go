package entities

import "time"

type Access struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:text"`
	Link      string `gorm:"type:text"`
	Priority  int64
	AccessID  *uint64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	SubAccess   *Access      `gorm:"foreignKey:AccessID"`
	Permissions []Permission `gorm:"foreignKey:AccessID"`
}

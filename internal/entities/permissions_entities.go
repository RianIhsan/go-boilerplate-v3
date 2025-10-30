package entities

import "time"

type Permission struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(100)"`
	Path      string `gorm:"type:varchar(100)"`
	Method    string `gorm:"type:varchar(100)"`
	AccessID  *uint64
	Type      string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Access *Access `gorm:"foreignKey:AccessID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

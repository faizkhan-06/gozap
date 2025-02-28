package models

import (
	"time"

	"gorm.io/gorm"
)

type Urls struct {
	ID        uint   `gorm:"primarykey;autoIncrement" json:"id"`
	ShortUrl  string `gorm:"not null" json:"short_url"`
	LongUrl   string `gorm:"not null" json:"long_url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
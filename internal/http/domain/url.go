package domain

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	OriginalURL string         `gorm:"not null" json:"original_url"`
	ShortCode   string         `gorm:"uniqueIndex;not null" json:"short_code"`
	ClickCount  int            `gorm:"default:0" json:"click_count"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

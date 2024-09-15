package models

import (
	"gorm.io/gorm"
	"time"
)

type NewsTopicOfTheSeasons struct {
	// @swaggerignore
	CreatedAt time.Time `json:"created_at"`
	// @swaggerignore
	UpdatedAt time.Time `json:"updated_at"`
	// @swaggerignore
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	NewsID             int            `gorm:"primaryKey;not null" json:"news-id"`
	TopicOfTheSeasonID int            `gorm:"primaryKey;not null" json:"topic-of-the-season-id"`
}

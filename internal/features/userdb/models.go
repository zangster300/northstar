package userdb

import (
	"time"

	"gorm.io/gorm"
)

type DiaryEntry struct {
	gorm.Model
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

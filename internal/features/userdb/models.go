package userdb

import (
	"time"

	"gorm.io/gorm"
)

type DiaryEntry struct {
	gorm.Model
	UserUUID  string    `json:"user_uuid" gorm:"index;not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
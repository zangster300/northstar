package auth

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	bcryptCost = 14
)

type User struct {
	gorm.Model
	UUID         string `json:"uuid" gorm:"uniqueIndex;not null"`
	Username     string `json:"username" gorm:"uniqueIndex;not null;size:255"`
	PasswordHash string `json:"-" gorm:"not null"`
}

var db *gorm.DB

func InitDB(dbPath string) error {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return err
	}

	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}

	slog.Info("Database initialized successfully", slog.String("path", dbPath))
	return nil
}

func CreateUser(username, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		UUID:         uuid.New().String(),
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUUID(userUUID string) (*User, error) {
	var user User
	if err := db.Where("uuid = ?", userUUID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func ValidatePassword(user *User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

func UserExists(username string) (bool, error) {
	var count int64
	if err := db.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

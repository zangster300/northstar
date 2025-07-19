package userdb

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitUserDatabase(userUUID string) error {
	dbPath := GetUserDatabasePath(userUUID)

	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return fmt.Errorf("failed to create users directory: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to create user database %s: %w", dbPath, err)
	}

	if err := db.AutoMigrate(&DiaryEntry{}); err != nil {
		return fmt.Errorf("failed to migrate user database %s: %w", dbPath, err)
	}

	if err := CloseUserDatabase(db); err != nil {
		slog.Warn("Failed to close user database connection", slog.String("path", dbPath), slog.Any("error", err))
	}

	slog.Info("User database initialized successfully",
		slog.String("user_uuid", userUUID),
		slog.String("path", dbPath))

	return nil
}

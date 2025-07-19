package userdb

import (
	"fmt"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetUserDatabasePath(userUUID string) string {
	return filepath.Join("data", "users", fmt.Sprintf("%s.db", userUUID))
}

func ConnectToUserDatabase(userUUID string) (*gorm.DB, error) {
	dbPath := GetUserDatabasePath(userUUID)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user database %s: %w", dbPath, err)
	}
	return db, nil
}

func CloseUserDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying database: %w", err)
	}
	return sqlDB.Close()
}

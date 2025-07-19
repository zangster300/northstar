package diary

import (
	"fmt"
	"time"

	"northstar/internal/features/userdb"
)

func CreateDiaryEntry(userUUID, content string) error {
	db, err := userdb.ConnectToUserDatabase(userUUID)
	if err != nil {
		return fmt.Errorf("failed to connect to user database: %w", err)
	}
	defer userdb.CloseUserDatabase(db)

	entry := &userdb.DiaryEntry{
		UserUUID:  userUUID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := db.Create(entry).Error; err != nil {
		return fmt.Errorf("failed to create diary entry: %w", err)
	}

	return nil
}

func GetDiaryEntries(userUUID string) ([]userdb.DiaryEntry, error) {
	db, err := userdb.ConnectToUserDatabase(userUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user database: %w", err)
	}
	defer userdb.CloseUserDatabase(db)

	var entries []userdb.DiaryEntry
	if err := db.Where("user_uuid = ?", userUUID).Order("created_at DESC").Find(&entries).Error; err != nil {
		return nil, fmt.Errorf("failed to get diary entries: %w", err)
	}

	return entries, nil
}

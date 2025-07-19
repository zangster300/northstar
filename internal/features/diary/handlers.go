package diary

import (
	"log/slog"
	"net/http"
	"strings"

	"northstar/internal/features/auth"
	"northstar/internal/features/diary/pages"

	"github.com/starfederation/datastar-go/datastar"
)

func HandleDiaryPage(w http.ResponseWriter, r *http.Request) {
	user, ok := auth.GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	entries, err := GetDiaryEntries(user.UUID)
	if err != nil {
		slog.Error("Failed to get diary entries", slog.Any("error", err), slog.String("user_uuid", user.UUID))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := pages.DiaryPage(entries).Render(r.Context(), w); err != nil {
		slog.Error("Failed to render diary page", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func HandleDiarySubmit(w http.ResponseWriter, r *http.Request) {
	user, ok := auth.GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	content := strings.TrimSpace(r.FormValue("content"))
	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	if err := CreateDiaryEntry(user.UUID, content); err != nil {
		slog.Error("Failed to create diary entry", slog.Any("error", err), slog.String("user_uuid", user.UUID))
		http.Error(w, "Failed to save diary entry", http.StatusInternalServerError)
		return
	}

	entries, err := GetDiaryEntries(user.UUID)
	if err != nil {
		slog.Error("Failed to get updated diary entries", slog.Any("error", err), slog.String("user_uuid", user.UUID))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	if err := sse.PatchElementTempl(pages.DiaryContent(entries)); err != nil {
		slog.Error("Failed to patch diary content fragment", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

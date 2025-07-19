package auth

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"gorm.io/gorm"
)

type contextKey string

const (
	userContextKey contextKey = "user"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userUUID, authenticated := GetUserFromSession(r)
		if !authenticated {
			http.Redirect(w, r, "/auth/login?redirect="+r.URL.Path, http.StatusSeeOther)
			return
		}

		user, err := GetUserByUUID(userUUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Warn("User not found in database", slog.String("uuid", userUUID))
			} else {
				slog.Error("Failed to get user from database", slog.Any("error", err))
			}
			ClearSession(w, r)
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RedirectIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userUUID, authenticated := GetUserFromSession(r)
		if authenticated {
			_, err := GetUserByUUID(userUUID)
			if err == nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			ClearSession(w, r)
		}
		next.ServeHTTP(w, r)
	})
}

func GetUserFromContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(userContextKey).(*User)
	return user, ok
}

type RateLimiter struct {
	attempts map[string][]time.Time
	mu       sync.RWMutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		attempts: make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) Allow(clientIP string, maxAttempts int, window time.Duration) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-window)

	if attempts, exists := rl.attempts[clientIP]; exists {
		validAttempts := make([]time.Time, 0, len(attempts))
		for _, attempt := range attempts {
			if attempt.After(cutoff) {
				validAttempts = append(validAttempts, attempt)
			}
		}
		rl.attempts[clientIP] = validAttempts

		if len(validAttempts) >= maxAttempts {
			return false
		}
	}

	rl.attempts[clientIP] = append(rl.attempts[clientIP], now)
	return true
}

func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cutoff := time.Now().Add(-time.Hour)
	for clientIP, attempts := range rl.attempts {
		validAttempts := make([]time.Time, 0, len(attempts))
		for _, attempt := range attempts {
			if attempt.After(cutoff) {
				validAttempts = append(validAttempts, attempt)
			}
		}
		if len(validAttempts) == 0 {
			delete(rl.attempts, clientIP)
		} else {
			rl.attempts[clientIP] = validAttempts
		}
	}
}

func StartRateLimiterCleanup(rl *RateLimiter) {
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			rl.cleanup()
		}
	}()
}
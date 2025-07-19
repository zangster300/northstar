package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"northstar/internal/features/auth/pages"
	"gorm.io/gorm"
)

var loginRateLimiter = NewRateLimiter()

func init() {
	StartRateLimiterCleanup(loginRateLimiter)
}

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	redirect := r.URL.Query().Get("redirect")
	if redirect == "" {
		redirect = "/"
	}

	if err := pages.LoginPage(redirect, "").Render(r.Context(), w); err != nil {
		slog.Error("Failed to render login page", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientIP := getClientIP(r)
	if !loginRateLimiter.Allow(clientIP, 5, 15*time.Minute) {
		slog.Warn("Rate limit exceeded for login", slog.String("client_ip", clientIP))
		renderLoginWithError(w, r, "Too many login attempts. Please try again later.")
		return
	}

	username := strings.TrimSpace(r.FormValue("username"))
	password := r.FormValue("password")
	redirect := r.FormValue("redirect")

	if username == "" || password == "" {
		renderLoginWithError(w, r, "Username and password are required")
		return
	}

	user, err := GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			renderLoginWithError(w, r, "Invalid username or password")
			return
		}
		slog.Error("Database error during login", slog.Any("error", err))
		renderLoginWithError(w, r, "Internal server error")
		return
	}

	if !ValidatePassword(user, password) {
		renderLoginWithError(w, r, "Invalid username or password")
		return
	}

	if err := SetUserSession(w, r, user.UUID); err != nil {
		slog.Error("Failed to set user session", slog.Any("error", err))
		renderLoginWithError(w, r, "Internal server error")
		return
	}

	if redirect == "" || !isValidRedirect(redirect) {
		redirect = "/"
	}

	slog.Info("User logged in", slog.String("username", username))
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}

func HandleSignupPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.SignupPage("").Render(r.Context(), w); err != nil {
		slog.Error("Failed to render signup page", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	clientIP := getClientIP(r)
	if !loginRateLimiter.Allow(clientIP, 3, 15*time.Minute) {
		slog.Warn("Rate limit exceeded for signup", slog.String("client_ip", clientIP))
		renderSignupWithError(w, r, "Too many signup attempts. Please try again later.")
		return
	}

	username := strings.TrimSpace(r.FormValue("username"))
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if username == "" || password == "" || confirmPassword == "" {
		renderSignupWithError(w, r, "All fields are required")
		return
	}

	if len(username) < 3 {
		renderSignupWithError(w, r, "Username must be at least 3 characters long")
		return
	}

	if len(password) < 8 {
		renderSignupWithError(w, r, "Password must be at least 8 characters long")
		return
	}

	if password != confirmPassword {
		renderSignupWithError(w, r, "Passwords do not match")
		return
	}

	if !isValidUsername(username) {
		renderSignupWithError(w, r, "Username can only contain letters, numbers, and underscores")
		return
	}

	exists, err := UserExists(username)
	if err != nil {
		slog.Error("Database error during signup", slog.Any("error", err))
		renderSignupWithError(w, r, "Internal server error")
		return
	}

	if exists {
		renderSignupWithError(w, r, "Username already exists")
		return
	}

	user, err := CreateUser(username, password)
	if err != nil {
		slog.Error("Failed to create user", slog.Any("error", err))
		renderSignupWithError(w, r, "Internal server error")
		return
	}

	if err := SetUserSession(w, r, user.UUID); err != nil {
		slog.Error("Failed to set user session after signup", slog.Any("error", err))
		renderSignupWithError(w, r, "Account created but login failed. Please try logging in.")
		return
	}

	slog.Info("New user created", slog.String("username", username))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	if err := ClearSession(w, r); err != nil {
		slog.Error("Failed to clear session during logout", slog.Any("error", err))
	}
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func renderLoginWithError(w http.ResponseWriter, r *http.Request, errorMsg string) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/"
	}

	if err := pages.LoginPage(redirect, errorMsg).Render(r.Context(), w); err != nil {
		slog.Error("Failed to render login page with error", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func renderSignupWithError(w http.ResponseWriter, r *http.Request, errorMsg string) {
	if err := pages.SignupPage(errorMsg).Render(r.Context(), w); err != nil {
		slog.Error("Failed to render signup page with error", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func getClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	xri := r.Header.Get("X-Real-Ip")
	if xri != "" {
		return strings.TrimSpace(xri)
	}

	return strings.Split(r.RemoteAddr, ":")[0]
}

func isValidRedirect(redirect string) bool {
	return strings.HasPrefix(redirect, "/") && !strings.HasPrefix(redirect, "//")
}

func isValidUsername(username string) bool {
	for _, char := range username {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '_') {
			return false
		}
	}
	return true
}

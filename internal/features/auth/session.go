package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func InitSessionStore() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   false,
	}
}

func SetUserSession(w http.ResponseWriter, r *http.Request, userUUID string) error {
	session, _ := store.Get(r, "northstar-session")
	session.Values["user_uuid"] = userUUID
	return session.Save(r, w)
}

func GetUserFromSession(r *http.Request) (string, bool) {
	session, _ := store.Get(r, "northstar-session")
	userUUID, ok := session.Values["user_uuid"].(string)
	return userUUID, ok && userUUID != ""
}

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "northstar-session")
	session.Values["user_uuid"] = nil
	session.Options.MaxAge = -1
	return session.Save(r, w)
}
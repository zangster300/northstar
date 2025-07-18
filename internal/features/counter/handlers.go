package counter

import (
	"net/http"
	"sync/atomic"

	"northstar/internal/features/counter/pages"

	"github.com/Jeffail/gabs/v2"
	"github.com/gorilla/sessions"
	"github.com/starfederation/datastar-go/datastar"
)

var globalCounter atomic.Uint32

const (
	sessionKey = "counter"
	countKey   = "count"
)

func GetUserValue(sessionStore sessions.Store) func(r *http.Request) (uint32, *sessions.Session, error) {
	return func(r *http.Request) (uint32, *sessions.Session, error) {
		session, err := sessionStore.Get(r, sessionKey)
		if err != nil {
			return 0, nil, err
		}

		val, ok := session.Values[countKey].(uint32)
		if !ok {
			val = 0
		}
		return val, session, nil
	}
}

func HandleCounterPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.CounterInitial().Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandleCounterData(getUserValue func(r *http.Request) (uint32, *sessions.Session, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userCount, _, err := getUserValue(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		store := pages.CounterSignals{
			Global: globalCounter.Load(),
			User:   userCount,
		}

		if err := datastar.NewSSE(w, r).PatchElementTempl(pages.Counter(store)); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func updateGlobal(store *gabs.Container) {
	_, _ = store.Set(globalCounter.Add(1), "global")
}

func HandleIncrementGlobal(w http.ResponseWriter, r *http.Request) {
	update := gabs.New()
	updateGlobal(update)

	if err := datastar.NewSSE(w, r).MarshalAndPatchSignals(update); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandleIncrementUser(getUserValue func(r *http.Request) (uint32, *sessions.Session, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, sess, err := getUserValue(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		val++
		sess.Values[countKey] = val
		if err := sess.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		update := gabs.New()
		updateGlobal(update)
		if _, err := update.Set(val, "user"); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if err := datastar.NewSSE(w, r).MarshalAndPatchSignals(update); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

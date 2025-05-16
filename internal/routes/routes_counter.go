package routes

import (
	"net/http"
	"sync/atomic"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	datastar "github.com/starfederation/datastar/sdk/go"
	"github.com/zangster300/northstar/internal/ui/pages"
)

func setupCounterRoute(router chi.Router, sessionStore sessions.Store) error {

	router.Get("/counter", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.CounterInitial().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	var globalCounter atomic.Uint32
	const (
		sessionKey = "counter"
		countKey   = "count"
	)

	GetUserValue := func(r *http.Request) (uint32, *sessions.Session, error) {
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

	router.Get("/counter/data", func(w http.ResponseWriter, r *http.Request) {
		userCount, _, err := GetUserValue(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		store := pages.CounterSignals{
			Global: globalCounter.Load(),
			User:   userCount,
		}

		if err := datastar.NewSSE(w, r).MergeFragmentTempl(pages.Counter(store)); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	updateGlobal := func(store *gabs.Container) {
		_, _ = store.Set(globalCounter.Add(1), "global")
	}

	router.Route("/counter/increment", func(incrementRouter chi.Router) {
		incrementRouter.Post("/global", func(w http.ResponseWriter, r *http.Request) {
			update := gabs.New()
			updateGlobal(update)

			if err := datastar.NewSSE(w, r).MarshalAndMergeSignals(update); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		})

		incrementRouter.Post("/user", func(w http.ResponseWriter, r *http.Request) {
			val, sess, err := GetUserValue(r)
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

			if err := datastar.NewSSE(w, r).MarshalAndMergeSignals(update); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		})
	})

	return nil
}

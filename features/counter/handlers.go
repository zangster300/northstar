package counter

import (
	"net/http"
	"sync/atomic"

	"northstar/features/counter/pages"

	"github.com/Jeffail/gabs/v2"
	"github.com/gorilla/sessions"
	"github.com/starfederation/datastar-go/datastar"
)

const (
	sessionKey = "counter"
	countKey   = "count"
)

type Handlers struct {
	globalCounter atomic.Uint32
	sessionStore  sessions.Store
}

func NewHandlers(sessionStore sessions.Store) *Handlers {
	return &Handlers{
		sessionStore: sessionStore,
	}
}

func (h *Handlers) CounterPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.CounterPage().Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) CounterData(w http.ResponseWriter, r *http.Request) {
	userCount, _, err := h.getUserValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	store := pages.CounterSignals{
		Global: h.globalCounter.Load(),
		User:   userCount,
	}

	if err := datastar.NewSSE(w, r).PatchElementTempl(pages.Counter(store)); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) IncrementGlobal(w http.ResponseWriter, r *http.Request) {
	update := gabs.New()
	h.updateGlobal(update)

	if err := datastar.NewSSE(w, r).MarshalAndPatchSignals(update); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) IncrementUser(w http.ResponseWriter, r *http.Request) {
	val, sess, err := h.getUserValue(r)
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
	h.updateGlobal(update)
	if _, err := update.Set(val, "user"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := datastar.NewSSE(w, r).MarshalAndPatchSignals(update); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) getUserValue(r *http.Request) (uint32, *sessions.Session, error) {
	session, err := h.sessionStore.Get(r, sessionKey)
	if err != nil {
		return 0, nil, err
	}

	val, ok := session.Values[countKey].(uint32)
	if !ok {
		val = 0
	}
	return val, session, nil
}

func (h *Handlers) updateGlobal(store *gabs.Container) {
	_, _ = store.Set(h.globalCounter.Add(1), "global")
}

package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/delaneyj/toolbelt"
	datastar "github.com/starfederation/datastar-go/datastar"

	"northstar/internal/features/todo/pages"
	"northstar/internal/features/todo/pages/components"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

// In-memory storage for todo data
type TodoStore struct {
	mu   sync.RWMutex
	data map[string]*components.TodoMVC

	// Channel for broadcasting updates
	updateCh chan UpdateEvent
	// Map of session IDs to their update channels
	subscribers map[string]chan UpdateEvent
	subMu       sync.RWMutex
}

type UpdateEvent struct {
	SessionID string
	Data      *components.TodoMVC
}

var globalTodoStore = &TodoStore{
	data:        make(map[string]*components.TodoMVC),
	updateCh:    make(chan UpdateEvent, 100),
	subscribers: make(map[string]chan UpdateEvent),
}

func init() {
	// Start the update broadcaster goroutine
	go globalTodoStore.broadcaster()
}

func (ts *TodoStore) broadcaster() {
	for update := range ts.updateCh {
		ts.subMu.RLock()
		for sessionID, ch := range ts.subscribers {
			if sessionID == update.SessionID {
				select {
				case ch <- update:
				default:
					// Channel is full, skip
				}
			}
		}
		ts.subMu.RUnlock()
	}
}

func (ts *TodoStore) Subscribe(sessionID string) chan UpdateEvent {
	ts.subMu.Lock()
	defer ts.subMu.Unlock()

	ch := make(chan UpdateEvent, 10)
	ts.subscribers[sessionID] = ch
	return ch
}

func (ts *TodoStore) Unsubscribe(sessionID string) {
	ts.subMu.Lock()
	defer ts.subMu.Unlock()

	if ch, exists := ts.subscribers[sessionID]; exists {
		close(ch)
		delete(ts.subscribers, sessionID)
	}
}

func (ts *TodoStore) Get(sessionID string) (*components.TodoMVC, bool) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	mvc, exists := ts.data[sessionID]
	return mvc, exists
}

func (ts *TodoStore) Set(sessionID string, mvc *components.TodoMVC) {
	ts.mu.Lock()
	ts.data[sessionID] = mvc
	ts.mu.Unlock()

	// Broadcast update
	select {
	case ts.updateCh <- UpdateEvent{SessionID: sessionID, Data: mvc}:
	default:
		// Channel is full, skip
	}
}

func setupIndexRoute(router chi.Router, store sessions.Store) error {
	resetMVC := func(mvc *components.TodoMVC) {
		mvc.Mode = components.TodoViewModeAll
		mvc.Todos = []*components.Todo{
			{Text: "Learn any backend language", Completed: true},
			{Text: "Learn Datastar", Completed: false},
			{Text: "Create Hypermedia", Completed: false},
			{Text: "???", Completed: false},
			{Text: "Profit", Completed: false},
		}
		mvc.EditingIdx = -1
	}

	mvcSession := func(w http.ResponseWriter, r *http.Request) (string, *components.TodoMVC, error) {
		sessionID, err := upsertSessionID(store, r, w)
		if err != nil {
			return "", nil, fmt.Errorf("failed to get session id: %w", err)
		}

		mvc, exists := globalTodoStore.Get(sessionID)
		if !exists {
			mvc = &components.TodoMVC{}
			resetMVC(mvc)
			globalTodoStore.Set(sessionID, mvc)
		}

		return sessionID, mvc, nil
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.Index("Northstar").Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/todos", func(todosRouter chi.Router) {
			todosRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {

				sessionID, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				sse := datastar.NewSSE(w, r)

				// Subscribe to updates for this session
				updateCh := globalTodoStore.Subscribe(sessionID)
				defer globalTodoStore.Unsubscribe(sessionID)

				// Send initial data
				c := components.TodosMVCView(mvc)
				if err := sse.PatchElementTempl(c); err != nil {
					if err := sse.ConsoleError(err); err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
					return
				}

				// Watch for updates
				ctx := r.Context()
				for {
					select {
					case <-ctx.Done():
						return
					case update := <-updateCh:
						if update.Data == nil {
							continue
						}
						c := components.TodosMVCView(update.Data)
						if err := sse.PatchElementTempl(c); err != nil {
							if err := sse.ConsoleError(err); err != nil {
								http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
							}
							return
						}
					}
				}
			})

			todosRouter.Put("/reset", func(w http.ResponseWriter, r *http.Request) {
				sessionID, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				resetMVC(mvc)
				globalTodoStore.Set(sessionID, mvc)
			})

			todosRouter.Put("/cancel", func(w http.ResponseWriter, r *http.Request) {

				sessionID, mvc, err := mvcSession(w, r)
				sse := datastar.NewSSE(w, r)
				if err != nil {
					if err := sse.ConsoleError(err); err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
					return
				}

				mvc.EditingIdx = -1
				globalTodoStore.Set(sessionID, mvc)
			})

			todosRouter.Put("/mode/{mode}", func(w http.ResponseWriter, r *http.Request) {

				sessionID, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				modeStr := chi.URLParam(r, "mode")
				modeRaw, err := strconv.Atoi(modeStr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				mode := components.TodoViewMode(modeRaw)
				if mode < components.TodoViewModeAll || mode > components.TodoViewModeCompleted {
					http.Error(w, "invalid mode", http.StatusBadRequest)
					return
				}

				mvc.Mode = mode
				globalTodoStore.Set(sessionID, mvc)
			})

			todosRouter.Route("/{idx}", func(todoRouter chi.Router) {
				routeIndex := func(w http.ResponseWriter, r *http.Request) (int, error) {
					idx := chi.URLParam(r, "idx")
					i, err := strconv.Atoi(idx)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return 0, err
					}
					return i, nil
				}

				todoRouter.Post("/toggle", func(w http.ResponseWriter, r *http.Request) {
					sessionID, mvc, err := mvcSession(w, r)

					sse := datastar.NewSSE(w, r)
					if err != nil {
						if err := sse.ConsoleError(err); err != nil {
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
						}
						return
					}

					i, err := routeIndex(w, r)
					if err != nil {
						if err := sse.ConsoleError(err); err != nil {
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
						}
						return
					}

					if i < 0 {
						setCompletedTo := false
						for _, todo := range mvc.Todos {
							if !todo.Completed {
								setCompletedTo = true
								break
							}
						}
						for _, todo := range mvc.Todos {
							todo.Completed = setCompletedTo
						}
					} else {
						todo := mvc.Todos[i]
						todo.Completed = !todo.Completed
					}

					globalTodoStore.Set(sessionID, mvc)
				})

				todoRouter.Route("/edit", func(editRouter chi.Router) {
					editRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
						sessionID, mvc, err := mvcSession(w, r)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}

						i, err := routeIndex(w, r)
						if err != nil {
							return
						}

						mvc.EditingIdx = i
						globalTodoStore.Set(sessionID, mvc)
					})

					editRouter.Put("/", func(w http.ResponseWriter, r *http.Request) {
						type Store struct {
							Input string `json:"input"`
						}
						store := &Store{}

						if err := datastar.ReadSignals(r, store); err != nil {
							http.Error(w, err.Error(), http.StatusBadRequest)
							return
						}

						if store.Input == "" {
							return
						}

						sessionID, mvc, err := mvcSession(w, r)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}

						i, err := routeIndex(w, r)
						if err != nil {
							return
						}

						if i >= 0 {
							mvc.Todos[i].Text = store.Input
						} else {
							mvc.Todos = append(mvc.Todos, &components.Todo{
								Text:      store.Input,
								Completed: false,
							})
						}
						mvc.EditingIdx = -1

						globalTodoStore.Set(sessionID, mvc)

					})
				})

				todoRouter.Delete("/", func(w http.ResponseWriter, r *http.Request) {
					i, err := routeIndex(w, r)
					if err != nil {
						return
					}

					sessionID, mvc, err := mvcSession(w, r)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					if i >= 0 {
						mvc.Todos = append(mvc.Todos[:i], mvc.Todos[i+1:]...)
					} else {
						// Filter out completed todos
						var filteredTodos []*components.Todo
						for _, todo := range mvc.Todos {
							if !todo.Completed {
								filteredTodos = append(filteredTodos, todo)
							}
						}
						mvc.Todos = filteredTodos
					}
					globalTodoStore.Set(sessionID, mvc)
				})
			})
		})
	})

	return nil
}

func MustJSONMarshal(v any) string {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func upsertSessionID(store sessions.Store, r *http.Request, w http.ResponseWriter) (string, error) {
	sess, err := store.Get(r, "connections")
	if err != nil {
		return "", fmt.Errorf("failed to get session: %w", err)
	}

	id, ok := sess.Values["id"].(string)

	if !ok {
		id = toolbelt.NextEncodedID()
		sess.Values["id"] = id
		if err := sess.Save(r, w); err != nil {
			return "", fmt.Errorf("failed to save session: %w", err)
		}
	}

	return id, nil
}

package routes

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/starfederation/datastar-go/datastar"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/samber/lo"
	"github.com/zangster300/northstar/internal/ui/components"
	"github.com/zangster300/northstar/internal/ui/pages"
)

var (
	// In-memory storage for todos per session
	todosStore      = make(map[string]*components.TodoMVC)
	todosStoreMutex sync.RWMutex
)

func setupIndexRoute(router chi.Router, store sessions.Store) error {
	saveMVC := func(sessionID string, mvc *components.TodoMVC) error {
		todosStoreMutex.Lock()
		defer todosStoreMutex.Unlock()
		todosStore[sessionID] = mvc
		return nil
	}

	loadMVC := func(sessionID string) (*components.TodoMVC, bool) {
		todosStoreMutex.RLock()
		defer todosStoreMutex.RUnlock()
		mvc, exists := todosStore[sessionID]
		return mvc, exists
	}

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

		mvc, exists := loadMVC(sessionID)
		if !exists {
			mvc = &components.TodoMVC{}
			resetMVC(mvc)
			if err := saveMVC(sessionID, mvc); err != nil {
				return "", nil, fmt.Errorf("failed to save mvc: %w", err)
			}
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
				_, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				sse := datastar.NewSSE(w, r)
				c := components.TodosMVCView(mvc)
				if err := sse.PatchElementTempl(c); err != nil {
					if err := sse.ConsoleError(err); err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
					return
				}
			})

			todosRouter.Put("/reset", func(w http.ResponseWriter, r *http.Request) {
				sessionID, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				resetMVC(mvc)
				if err := saveMVC(sessionID, mvc); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
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
				if err := saveMVC(sessionID, mvc); err != nil {
					if err := sse.ConsoleError(err); err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
					return
				}
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
				if err := saveMVC(sessionID, mvc); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
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

					if err := saveMVC(sessionID, mvc); err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
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
						if err := saveMVC(sessionID, mvc); err != nil {
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
						}
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

						if err := saveMVC(sessionID, mvc); err != nil {
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
						}
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
						mvc.Todos = lo.Filter(mvc.Todos, func(todo *components.Todo, i int) bool {
							return !todo.Completed
						})
					}
					if err := saveMVC(sessionID, mvc); err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
				})
			})
		})
	})

	return nil
}

func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}

func upsertSessionID(store sessions.Store, r *http.Request, w http.ResponseWriter) (string, error) {
	sess, err := store.Get(r, "connections")
	if err != nil {
		return "", fmt.Errorf("failed to get session: %w", err)
	}

	id, ok := sess.Values["id"].(string)

	if !ok {
		id = generateID()
		sess.Values["id"] = id
		if err := sess.Save(r, w); err != nil {
			return "", fmt.Errorf("failed to save session: %w", err)
		}
	}

	return id, nil
}
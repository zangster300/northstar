package index

import (
	"encoding/json"
	"net/http"
	"strconv"

	"northstar/features/index/components"
	"northstar/features/index/pages"
	"northstar/features/index/services"

	"github.com/go-chi/chi/v5"
	"github.com/starfederation/datastar-go/datastar"
)

type Handlers struct {
	todoService *services.TodoService
}

func NewHandlers(todoService *services.TodoService) *Handlers {
	return &Handlers{
		todoService: todoService,
	}
}

func (h *Handlers) IndexPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.IndexPage("Northstar").Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) TodosSSE(w http.ResponseWriter, r *http.Request) {
	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)

	// Watch for updates
	ctx := r.Context()
	watcher, err := h.todoService.WatchUpdates(ctx, sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer watcher.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case entry := <-watcher.Updates():
			if entry == nil {
				continue
			}
			if err := json.Unmarshal(entry.Value(), mvc); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			c := components.TodosMVCView(mvc)
			if err := sse.PatchElementTempl(c); err != nil {
				if err := sse.ConsoleError(err); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
				return
			}
		}
	}
}

func (h *Handlers) ResetTodos(w http.ResponseWriter, r *http.Request) {
	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.todoService.ResetMVC(mvc)
	if err := h.todoService.SaveMVC(r.Context(), sessionID, mvc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) CancelEdit(w http.ResponseWriter, r *http.Request) {
	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
	sse := datastar.NewSSE(w, r)
	if err != nil {
		if err := sse.ConsoleError(err); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.todoService.CancelEditing(mvc)
	if err := h.todoService.SaveMVC(r.Context(), sessionID, mvc); err != nil {
		if err := sse.ConsoleError(err); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
}

func (h *Handlers) SetMode(w http.ResponseWriter, r *http.Request) {
	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
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

	h.todoService.SetMode(mvc, mode)
	if err := h.todoService.SaveMVC(r.Context(), sessionID, mvc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
	sse := datastar.NewSSE(w, r)
	if err != nil {
		if err := sse.ConsoleError(err); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	i, err := h.parseIndex(w, r)
	if err != nil {
		if err := sse.ConsoleError(err); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.todoService.ToggleTodo(mvc, i)
	if err := h.todoService.SaveMVC(r.Context(), sessionID, mvc); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) StartEdit(w http.ResponseWriter, r *http.Request) {
	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	i, err := h.parseIndex(w, r)
	if err != nil {
		return
	}

	h.todoService.StartEditing(mvc, i)
	if err := h.todoService.SaveMVC(r.Context(), sessionID, mvc); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) SaveEdit(w http.ResponseWriter, r *http.Request) {
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

	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	i, err := h.parseIndex(w, r)
	if err != nil {
		return
	}

	h.todoService.EditTodo(mvc, i, store.Input)
	if err := h.todoService.SaveMVC(r.Context(), sessionID, mvc); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	i, err := h.parseIndex(w, r)
	if err != nil {
		return
	}

	sessionID, mvc, err := h.todoService.GetSessionMVC(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.todoService.DeleteTodo(mvc, i)
	if err := h.todoService.SaveMVC(r.Context(), sessionID, mvc); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) parseIndex(w http.ResponseWriter, r *http.Request) (int, error) {
	idx := chi.URLParam(r, "idx")
	i, err := strconv.Atoi(idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}
	return i, nil
}

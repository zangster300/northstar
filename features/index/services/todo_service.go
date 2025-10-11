package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"northstar/features/index/components"

	"github.com/delaneyj/toolbelt"
	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/gorilla/sessions"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/samber/lo"
)

type TodoService struct {
	kv    jetstream.KeyValue
	store sessions.Store
}

func NewTodoService(ns *embeddednats.Server, store sessions.Store) (*TodoService, error) {
	nc, err := ns.Client()
	if err != nil {
		return nil, fmt.Errorf("error creating nats client: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("error creating jetstream client: %w", err)
	}

	kv, err := js.CreateOrUpdateKeyValue(context.Background(), jetstream.KeyValueConfig{
		Bucket:      "todos",
		Description: "Datastar Todos",
		Compression: true,
		TTL:         time.Hour,
		MaxBytes:    16 * 1024 * 1024,
	})

	if err != nil {
		return nil, fmt.Errorf("error creating key value: %w", err)
	}

	return &TodoService{
		kv:    kv,
		store: store,
	}, nil
}

func (s *TodoService) GetSessionMVC(w http.ResponseWriter, r *http.Request) (string, *components.TodoMVC, error) {
	ctx := r.Context()
	sessionID, err := s.upsertSessionID(r, w)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get session id: %w", err)
	}

	mvc := &components.TodoMVC{}
	if entry, err := s.kv.Get(ctx, sessionID); err != nil {
		if err != jetstream.ErrKeyNotFound {
			return "", nil, fmt.Errorf("failed to get key value: %w", err)
		}
		s.resetMVC(mvc)

		if err := s.saveMVC(ctx, sessionID, mvc); err != nil {
			return "", nil, fmt.Errorf("failed to save mvc: %w", err)
		}
	} else {
		if err := json.Unmarshal(entry.Value(), mvc); err != nil {
			return "", nil, fmt.Errorf("failed to unmarshal mvc: %w", err)
		}
	}
	return sessionID, mvc, nil
}

func (s *TodoService) SaveMVC(ctx context.Context, sessionID string, mvc *components.TodoMVC) error {
	return s.saveMVC(ctx, sessionID, mvc)
}

func (s *TodoService) ResetMVC(mvc *components.TodoMVC) {
	s.resetMVC(mvc)
}

func (s *TodoService) WatchUpdates(ctx context.Context, sessionID string) (jetstream.KeyWatcher, error) {
	return s.kv.Watch(ctx, sessionID)
}

func (s *TodoService) ToggleTodo(mvc *components.TodoMVC, index int) {
	if index < 0 {
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
	} else if index < len(mvc.Todos) {
		todo := mvc.Todos[index]
		todo.Completed = !todo.Completed
	}
}

func (s *TodoService) EditTodo(mvc *components.TodoMVC, index int, text string) {
	if index >= 0 && index < len(mvc.Todos) {
		mvc.Todos[index].Text = text
	} else if index < 0 {
		mvc.Todos = append(mvc.Todos, &components.Todo{
			Text:      text,
			Completed: false,
		})
	}
	mvc.EditingIdx = -1
}

func (s *TodoService) DeleteTodo(mvc *components.TodoMVC, index int) {
	if index >= 0 && index < len(mvc.Todos) {
		mvc.Todos = append(mvc.Todos[:index], mvc.Todos[index+1:]...)
	} else if index < 0 {
		mvc.Todos = lo.Filter(mvc.Todos, func(todo *components.Todo, i int) bool {
			return !todo.Completed
		})
	}
}

func (s *TodoService) SetMode(mvc *components.TodoMVC, mode components.TodoViewMode) {
	mvc.Mode = mode
}

func (s *TodoService) StartEditing(mvc *components.TodoMVC, index int) {
	mvc.EditingIdx = index
}

func (s *TodoService) CancelEditing(mvc *components.TodoMVC) {
	mvc.EditingIdx = -1
}

func (s *TodoService) saveMVC(ctx context.Context, sessionID string, mvc *components.TodoMVC) error {
	b, err := json.Marshal(mvc)
	if err != nil {
		return fmt.Errorf("failed to marshal mvc: %w", err)
	}
	if _, err := s.kv.Put(ctx, sessionID, b); err != nil {
		return fmt.Errorf("failed to put key value: %w", err)
	}
	return nil
}

func (s *TodoService) resetMVC(mvc *components.TodoMVC) {
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

func (s *TodoService) upsertSessionID(r *http.Request, w http.ResponseWriter) (string, error) {
	sess, err := s.store.Get(r, "connections")
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

package usecase

import (
	"github.com/ars0915/glossika-exercise/repo"
)

type AppHandler struct {
	Task
	User
}

type NewHandlerOption func(*AppHandler)

func newHandler(optFn ...NewHandlerOption) *AppHandler {
	h := &AppHandler{}

	for _, o := range optFn {
		o(h)
	}

	return h
}

type TaskHandler struct {
	db repo.App
}

func NewTaskHandler(db repo.App) *TaskHandler {
	return &TaskHandler{
		db: db,
	}
}

func WithTask(i Task) func(h *AppHandler) {
	return func(h *AppHandler) { h.Task = i }
}

type UserHandler struct {
	db    repo.App
	email repo.Email
	redis repo.Redis
}

func NewUserHandler(db repo.App, redis repo.Redis, email repo.Email) *UserHandler {
	return &UserHandler{
		db:    db,
		email: email,
		redis: redis,
	}
}

func WithUser(i User) func(h *AppHandler) {
	return func(h *AppHandler) { h.User = i }
}

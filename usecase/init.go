package usecase

import "github.com/ars0915/glossika-exercise/repo"

func InitHandler(db repo.App, redis repo.Redis, email repo.Email) Handler {
	user := NewUserHandler(db, redis, email)
	task := NewTaskHandler(db)

	h := newHandler(
		WithTask(task),
		WithUser(user),
	)

	return h
}

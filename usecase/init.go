package usecase

import "github.com/ars0915/glossika-exercise/repo"

func InitHandler(db repo.App, redis repo.Redis) Handler {
	task := NewTaskHandler(db)

	h := newHandler(
		WithTask(task),
	)

	return h
}

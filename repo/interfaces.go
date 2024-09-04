package repo

import (
	"github.com/ars0915/glossika-exercise/entity"
)

//go:generate mockgen -destination=../mocks/repo/app_repo.go -package=mocks github.com/ars0915/glossika-exercise/repo App

type (
	App interface {
		Migrate()
		Debug()

		// transaction
		Begin() App
		Commit() error
		Rollback() error

		ListTasks(param entity.ListTaskParam) (t []entity.Task, err error)
		GetTasksCount() (count int64, err error)
		GetTask(id uint) (task entity.Task, err error)
		CreateTask(t entity.Task) (entity.Task, error)
		UpdateTask(id uint, t entity.Task) error
		DeleteTask(id uint) (err error)

		CreateUser(t entity.User) (entity.User, error)
		UpdateUser(id uint, t entity.User) error
		GetUser(email string) (User entity.User, err error)
		GetUserForUpdate(email string) (user entity.User, err error)
	}

	Redis interface {
	}

	Email interface {
		SendVerificationEmail(email, verificationCode string) error
	}
)

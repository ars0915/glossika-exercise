package usecase

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/ars0915/glossika-exercise/entity"
	"github.com/ars0915/glossika-exercise/repo"
)

type RegisterParam struct {
	Email    string
	Password string
}

func (h UserHandler) Register(ctx context.Context, param RegisterParam) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}

	verificationCode, err := GenerateRandomString(6)
	if err != nil {
		return errors.Wrap(err, "GenerateRandomString")
	}

	return repo.WithinTransaction(ctx, h.db, func(txCtx context.Context) error {
		tx := repo.ExtractTx(txCtx)

		user, err := tx.CreateUser(entity.User{
			Email:            param.Email,
			Password:         string(hashedPassword),
			EmailVerified:    false,
			VerificationCode: verificationCode,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return ErrorEmailRegistered
			}
			return errors.Wrap(err, "CreateUser")
		}

		if err = h.email.SendVerificationEmail(user.Email, user.VerificationCode); err != nil {
			return errors.Wrap(err, "SendVerificationEmail")
		}

		return nil
	})
}

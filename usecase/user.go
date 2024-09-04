package usecase

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/ars0915/glossika-exercise/entity"
	"github.com/ars0915/glossika-exercise/repo"
	"github.com/ars0915/glossika-exercise/util/cTypes"
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
			Email:            cTypes.String(param.Email),
			Password:         cTypes.String(string(hashedPassword)),
			EmailVerified:    cTypes.Bool(false),
			VerificationCode: cTypes.String(verificationCode),
		})
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return ErrorEmailRegistered
			}
			return errors.Wrap(err, "CreateUser")
		}

		if err = h.email.SendVerificationEmail(*user.Email, *user.VerificationCode); err != nil {
			return errors.Wrap(err, "SendVerificationEmail")
		}

		return nil
	})
}

type VerifyUserParam struct {
	Email            string
	VerificationCode string
}

func (h UserHandler) Verify(ctx context.Context, param VerifyUserParam) error {
	return repo.WithinTransaction(ctx, h.db, func(txCtx context.Context) error {
		tx := repo.ExtractTx(txCtx)

		user, err := tx.GetUserForUpdate(param.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrorUserNotFound
			}
			return errors.Wrap(err, "GetUserForUpdate")
		}

		if *user.VerificationCode != param.VerificationCode {
			return ErrorUserVerificationFailed
		}

		if err = tx.UpdateUser(user.ID, entity.User{
			EmailVerified: cTypes.Bool(true),
		}); err != nil {
			return errors.Wrap(err, "UpdateUser")
		}

		return nil
	})
}

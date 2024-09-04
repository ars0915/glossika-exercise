package db

import (
	"gorm.io/gorm/clause"

	"github.com/ars0915/glossika-exercise/entity"
)

func (s *AppRepo) CreateUser(t entity.User) (entity.User, error) {
	if err := s.db.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (s *AppRepo) UpdateUser(id uint, t entity.User) error {
	return s.db.Model(entity.User{}).Where(`"id" = ?`, id).Updates(t).Error
}

func (s *AppRepo) GetUser(email string) (User entity.User, err error) {
	if err = s.db.Where(`"email" = ?`, email).First(&User).Error; err != nil {
		return User, err
	}
	return
}

func (s *AppRepo) GetUserForUpdate(email string) (user entity.User, err error) {
	if err = s.db.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return
}

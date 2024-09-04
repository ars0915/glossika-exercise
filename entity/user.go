package entity

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint           `json:"id" gorm:"primary_key"`
	Email            string         `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password         string         `json:"password" gorm:"type:varchar(60);not null"`
	EmailVerified    bool           `json:"emailVerified" gorm:"type:boolean;default:false"`
	VerificationCode string         `json:"verificationCode" gorm:"type:varchar(6);not null"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}

func (t User) MarshalJSON() ([]byte, error) {
	type Alias User
	s := struct {
		*Alias
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
		DeletedAt *int64 `json:"deleted_at,omitempty"`
	}{
		Alias:     (*Alias)(&t),
		CreatedAt: t.CreatedAt.Unix(),
		UpdatedAt: t.UpdatedAt.Unix(),
	}

	if t.DeletedAt.Valid {
		deletedAt := t.DeletedAt.Time.Unix()
		s.DeletedAt = &deletedAt
	}

	return json.Marshal(&s)
}

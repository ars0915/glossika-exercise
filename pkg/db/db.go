package db

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ars0915/glossika-exercise/config"
)

func NewDB(config config.ConfENV) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, errors.Wrap(err, "init db")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "connect db")
	}

	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping db")
	}

	return db, nil
}

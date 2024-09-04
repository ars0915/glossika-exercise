package db

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ars0915/glossika-exercise/config"
)

func NewDB(config config.ConfENV) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)

	for i := 0; i < config.DB.MaxConnectionRetry; i++ {
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			err = errors.Wrap(err, "init db")
			time.Sleep(config.DB.RetryDelay)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			err = errors.Wrap(err, "connect db")
			time.Sleep(config.DB.RetryDelay)
			continue
		}

		sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Hour)

		err = sqlDB.Ping()
		if err != nil {
			err = errors.Wrap(err, "ping db")
			time.Sleep(config.DB.RetryDelay)
			continue
		}
	}

	return db, err
}

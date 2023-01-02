package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/s-beats/go-cms/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewDB() (*DB, error) {
	dialector := newDialector()
	logger := newLogger()

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	return &DB{db}, nil
}

// FIXME: DSNの設定
func newDialector() gorm.Dialector {
	return mysql.Open(
		fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?parseTime=true",
			env.MySQLUser(),
			env.MySQLPassword(),
			env.MySQLHost(),
			env.MySQLPort(),
			env.MySQLDatabase(),
		),
	)
}

func newLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}

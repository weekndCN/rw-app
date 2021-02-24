package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB struct
type DB struct {
	Conn *gorm.DB
}

// Open connect to a database and ping db
func Open(driver, dsn string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// generate sql but no execute
		DryRun: false,
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	// sql pool
	sqlDB, err := db.DB()
	sqlDB.Ping()

	if err != nil {
		return nil, err
	}
	// max idle connection numbers
	sqlDB.SetMaxIdleConns(10)
	// max open connection numbers
	sqlDB.SetMaxOpenConns(100)
	// max life time
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &DB{Conn: db}, nil
}

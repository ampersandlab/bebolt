package internal

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ampersandlab/bebolt/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *Database
	dbOnce sync.Once
)

type Database struct {
	DB *gorm.DB
}

func NewDatabaseConnection(env *config.Env) *Database {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBDatabase)

	dbOnce.Do(func() {
		log.Println("CONNECTING TO: " + dns)
		conn, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dns,   // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{})
		// conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

		if err != nil {
			panic("Unable to set db connection.")
		}
		sqlDB, err := conn.DB()
		if err != nil {
			panic("Unable to get database")
		}
		sqlDB.SetMaxIdleConns(25)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)
		sqlDB.SetConnMaxLifetime(2 * time.Hour)
		db = &Database{
			DB: nil,
		}
	})

	return db
}

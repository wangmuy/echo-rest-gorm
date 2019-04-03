package db

import (
	"echo-rest-gorm/config"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init(c *config.Config) {
	var err error
	db, err = gorm.Open(c.DB_DIALECT, c.DB_CONN)
	if err != nil {
		fmt.Println(err)
		panic("DB Connection Error")
	}
	db.DB().SetMaxIdleConns(c.DB_MAX_IDLE_CONNS)
	db.DB().SetMaxOpenConns(c.DB_MAX_OPEN_CONNS)
	db.SingularTable(true) // singular lowercase table name
}

func DBManager() *gorm.DB {
	return db
}

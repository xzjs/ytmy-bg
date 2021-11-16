package lib

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbonce sync.Once

func GetDB() *gorm.DB {
	dbonce.Do(func() {
		c := GetConfig()
		dbconf := c.DB
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.User, dbconf.Pass, dbconf.Host, dbconf.Port, dbconf.DBName)
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}

package mysql

import (
	"log"

	"github.com/iandjx/hackernews/internal/links"
	"github.com/iandjx/hackernews/internal/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	dsn := "root:password@tcp(localhost:3306)/gorm_hackernews?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&users.User{}, &links.Link{})
	Db = db
}

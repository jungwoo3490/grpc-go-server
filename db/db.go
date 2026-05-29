package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
	Done        bool
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB 연결 실패: ", err)
	}

	if err := db.AutoMigrate(&Todo{}); err != nil {
		log.Fatal("마이그레이션 실패: ", err)
	}

	return db
}

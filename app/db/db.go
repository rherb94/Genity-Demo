package db

import (
	"fmt"
	"Genity-Demo/app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "genity.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("connection to database failed...")
	}
	defer db.Close()
	db.AutoMigrate(&entity.Data{})
}
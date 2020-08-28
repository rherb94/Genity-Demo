package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"Genity-Demo/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
func GetAllData(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "genity.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var data []entity.Data
	db.Find(&data)
	fmt.Println("{}", data)
	json.NewEncoder(w).Encode(data)
}

func PostData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Endpoint Hit!")
	db, err := gorm.Open("sqlite3", "genity.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	title := vars["title"]
	uuid4 := entity.GenerateID()
	timeStamp := entity.GetCurrentTime()


	fmt.Println("Title:", title)
	db.Create(&entity.Data{Title: title, UUID4: uuid4, TimeStamp: timeStamp})
	fmt.Fprintf(w, "Post Successful!")
}
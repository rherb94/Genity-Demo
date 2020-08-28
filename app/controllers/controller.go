package controller

import (
	"Genity-Demo/app/service"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/post-data/{title}", service.PostData).Methods("POST")
	myRouter.HandleFunc("/get-data", service.GetAllData).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

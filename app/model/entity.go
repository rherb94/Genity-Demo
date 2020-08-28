package entity

import (
	"time"
	"github.com/google/uuid"
	// "github.com/joeshaw/iso8601"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Data struct { 
	Title string 
	UUID4 string `json:"id"`
	TimeStamp string 

}

func GenerateID() string {
	return uuid.New().String()
}
func GetCurrentTime() string { 
	return time.Now().UTC().Format("2006-01-02T15:04:05-0700")
}
package main
import (
	"Genity-Demo/app/controllers"
	"Genity-Demo/app/db"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	fmt.Println("Herb Genity Demo")
	db.InitialMigration()
	controller.HandleRequests()
}


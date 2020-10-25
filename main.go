package main

import (
	"mojito/database"
	"mojito/model"
	"mojito/redis"
	"mojito/router"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var err error

func main() {

	// Creating a connection to the database
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer database.DB.Close()

	// run the migrations: todo struct
	database.DB.AutoMigrate(&model.Lost{})

	//setup routes 
	r := router.SetupRouter()

	// running
	r.Run()

	for {
		time.Sleep(2 * time.Second)
		taskValue := redis.RB.ReceiveTask("CreateLost")
		if len(taskValue) > 0 {
			go redis.HandleTask("CreateLost", taskValue)
		}
	}
}
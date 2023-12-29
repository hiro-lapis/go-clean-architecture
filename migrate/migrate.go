package main

import (
	"clean-architecture/db"
	"clean-architecture/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully connection")
	defer db.CloseDB(dbConn)
	// pass models that you want to migrate
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}

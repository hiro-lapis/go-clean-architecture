package main

import (
	"clean-architecture/db"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully connection")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate()
}

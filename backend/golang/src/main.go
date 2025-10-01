package main

import (
	"gymlink/internal/dbase"
	"log"
)

func main() {
	db := dbase.ConnectDB()
	defer db.Close()

	err := dbase.MigrateUp(db)
	if err != nil {
		log.Fatal("migrate up failed")
	}
}

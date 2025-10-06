package main

import (
	"gymlink/internal/dbase"
	"gymlink/internal/handler"
	"gymlink/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := dbase.ConnectDB()
	defer db.Close()

	err := dbase.MigrateUp(db)
	if err != nil {
		log.Fatal("migrate up failed")
	}

	err = dbase.SeedingDB(db)
	if err != nil {
		log.Fatal("seeding error")
	}

	userRepo := dbase.NewUserQueryRepo(db)

	userSvc, err := service.NewUserService(userRepo)
	if err != nil {
		log.Fatal("user service error")
	}

	h := handler.NewUserHandler(userSvc)

	r := gin.Default()
	r.POST("/users", h.GetUserByParam)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

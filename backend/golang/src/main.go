package main

import (
	"context"
	"gymlink/internal/adapter"
	"gymlink/internal/dbase"
	"gymlink/internal/handler"
	"gymlink/internal/service"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	// db へ接続関連
	db := dbase.ConnectDB()

	defer db.Close()

	//マイグレーション
	err := dbase.MigrateUp(db)
	if err != nil {
		log.Fatal("migrate up failed")
	}
	//seeding
	err = dbase.SeedingDB(db)
	if err != nil {
		log.Fatal("seeding error")
	}
	// dbase のコンストラクタ => userRepo が db とのやり取りのインターフェースのイメージ
	err = godotenv.Load(".env")
	if err != nil {
		wd, _ := os.Getwd()
		log.Fatal("error: no env file is found at base ", wd)
	}
	credPath := os.Getenv("FIREBASE_CREDENTIAL_PATH")
	if credPath == "" {
		log.Fatal("FIREBASE CREDENTIAL PATH is not set")
	}
	// firebase 関連
	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializing app : Firebase Error")
	}

	ctx := context.Background()
	authCli, err := app.Auth(ctx)
	if err != nil {
		log.Fatal("error in Auth")
	}
	authC := adapter.NewAuthClient(authCli)

	userRepo := dbase.NewUserQueryRepo(db)

	userSvc, err := service.NewUserService(userRepo, authC)
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

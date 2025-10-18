package main

import (
	"context"
	"fmt"
	"gymlink/internal/adapter"
	"gymlink/internal/dbase"
	"gymlink/internal/handler"
	"gymlink/internal/service"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {

	var client = &http.Client{}

	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	fmt.Println("region:", awsCfg.Region)
	if err != nil {
		log.Fatal("aws configuration error:", err)
	}
	// db へ接続関連
	db := dbase.ConnectDB()

	defer db.Close()

	//マイグレーション
	err = dbase.MigrateUp(db)
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

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: ", err)
	}
	awsCli := adapter.NewAwsClient(awsCfg, "katazuke")
	// awsCli.CheckBucket("katazuke")
	// awsCli.UploadImage(ctx, "katazuke", "myKey.png")

	apiKey := os.Getenv("GPT_API_KEY")
	baseUrl := os.Getenv("GPT_URL")
	if apiKey == "" || baseUrl == "" {
		log.Fatal("error api setting ", err)
	}

	gptCli := adapter.NewGptClient(client, apiKey, baseUrl)

	userQueryRepo := dbase.NewUserQueryRepo(db)

	userCreateRepo := dbase.NewUserCreateRepo(db)

	profileRepo := dbase.NewProfileRepo(db)

	exerciseQueryRepo := dbase.NewExerciseQueryRepo(db)

	exerciseCreateRepo := dbase.NewExerciseCreateRepo(db)

	userSvc, err := service.NewUserService(userQueryRepo, userCreateRepo, profileRepo, authC)
	if err != nil {
		log.Fatal("user service error")
	}

	exerciseSvc, err := service.NewExerciseService(exerciseQueryRepo, exerciseCreateRepo, authC, gptCli, awsCli)
	if err != nil {
		log.Fatal("exercise service error")
	}

	userHandler := handler.NewUserHandler(userSvc)
	exerciseHandler := handler.NewExerciseHandler(exerciseSvc)

	r := gin.Default()
	r.POST("/users", userHandler.SignUpUserById)
	r.POST("/login", userHandler.LoginUser)
	r.GET("/user_profiles/:user_id", userHandler.GetProfilebyId)
	r.GET("/users/:user_id/exercises", exerciseHandler.GetExercisesById)
	r.GET("/records", exerciseHandler.GetRecords)
	// r.POST("/records", exerciseHandler.GenerateIllustration)
	r.POST("/likes", exerciseHandler.CreateLike)
	r.GET("/likes/:record_id", exerciseHandler.CheckLike)
	r.DELETE("/likes/:record_id", exerciseHandler.DeleteLike)
	r.POST("/follows", userHandler.FollowUser)
	r.DELETE("/users/unfollows", userHandler.DeleteFollowUser)
	r.POST("/upload", exerciseHandler.GenerateIllustration)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

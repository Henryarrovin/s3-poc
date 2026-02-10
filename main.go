package main

import (
	"log"
	"s3-poc/config"
	"s3-poc/data"
	"s3-poc/handlers"
	"s3-poc/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	s3Client, err := config.NewS3Client()
	if err != nil {
		log.Fatal("failed to create S3 client:", err)
	}

	userRepo := data.NewUserRepository(db)
	s3Repo := data.NewS3Repository(s3Client)

	userHandler := handlers.NewUserHandler(userRepo, s3Repo)
	s3Handler := handlers.NewS3Handler(s3Repo)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/buckets/:name", s3Handler.CreateBucket)
	r.GET("/buckets/:name/objects", s3Handler.ListObjects)

	log.Println("Server started at :8080")
	r.Run(":8080")
}

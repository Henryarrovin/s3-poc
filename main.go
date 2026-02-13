package main

import (
	"log"
	"s3-poc/wire"

	"github.com/gin-gonic/gin"
)

func main() {

	userHandler, err := wire.InitializeUserHandler()
	if err != nil {
		log.Fatal(err)
	}

	s3Handler, err := wire.InitializeS3Handler()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)

	r.POST("/buckets/:name", s3Handler.CreateBucket)
	r.GET("/buckets/:name/objects", s3Handler.ListObjects)

	log.Println("Server running at :8080")
	r.Run(":8080")
}

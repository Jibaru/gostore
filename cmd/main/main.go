package main

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	address := ":80"
	bucketRepository := repositories.NewRamBucketRepository()
	objectRepository := repositories.NewRamObjectRepository()

	createBucketServ := application.NewCreateBucketService(bucketRepository)
	createObjectServ := application.NewCreateObjectService(
		bucketRepository,
		objectRepository,
	)

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/buckets", controllers.NewCreateBucket(createBucketServ).Handle)
	e.POST("/buckets/:bucketID/objects", controllers.NewCreateObject(createObjectServ).Handle)
	e.GET("/buckets", controllers.NewGetBuckets(bucketRepository).Handle)

	e.Logger.Fatal(e.Start(address))
}

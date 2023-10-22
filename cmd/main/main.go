package main

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/jibaru/gostore/internal/shared"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func main() {
	host := "localhost"
	port := 80
	address := host + ":" + strconv.Itoa(port)
	storageFolderName := "storage"

	urlGenerator := shared.NewUrlGenerator(
		"localhost",
		uint(port),
		storageFolderName,
		false,
	)
	filesystem := shared.NewServerFilesystem("./" + storageFolderName)

	bucketRepository, err := repositories.NewFileBucketRepository(storageFolderName + "/buckets.json")
	if err != nil {
		panic(err)
		return
	}
	objectRepository, err := repositories.NewFileObjectRepository(storageFolderName + "/objects.json")
	if err != nil {
		panic(err)
		return
	}

	getBucketsService := application.NewGetBucketsService(
		bucketRepository,
	)
	generateBucketPathServ := application.NewGenerateBucketPathService(
		bucketRepository,
	)
	generateObjectPathServ := application.NewGenerateObjectPathService(
		objectRepository,
		generateBucketPathServ,
	)
	createBucketServ := application.NewCreateBucketService(
		bucketRepository,
		filesystem,
		generateBucketPathServ,
	)
	createObjectServ := application.NewCreateObjectService(
		bucketRepository,
		objectRepository,
		filesystem,
		generateBucketPathServ,
	)
	getBucketsInBucketServ := application.NewGetBucketsInBucketService(
		bucketRepository,
	)
	getObjectsInBucketServ := application.NewGetObjectsInBucketService(
		objectRepository,
	)

	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/storage", storageFolderName)
	e.POST("/buckets", controllers.NewCreateBucket(createBucketServ).Handle)
	e.POST("/buckets/:bucketID/objects", controllers.NewCreateObject(createObjectServ).Handle)
	e.GET("/buckets/:bucketID/objects", controllers.NewGetObjectsInBucket(getObjectsInBucketServ).Handle)
	e.GET("/buckets", controllers.NewGetBuckets(getBucketsService).Handle)
	e.GET("/buckets/:bucketID/buckets", controllers.NewGetBucketsInBucket(getBucketsInBucketServ).Handle)
	e.GET("/objects/:objectID/download", controllers.NewDownloadObject(urlGenerator, generateObjectPathServ).Handle)

	e.Logger.Fatal(e.Start(address))
}

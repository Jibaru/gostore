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
	deleteObjectServ := application.NewDeleteObjectService(
		objectRepository,
		generateObjectPathServ,
		filesystem,
	)

	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/storage", storageFolderName)

	api := e.Group("/api/v1")
	api.POST("/buckets", controllers.NewCreateBucket(createBucketServ).Handle)
	api.POST("/buckets/:bucketID/objects", controllers.NewCreateObject(createObjectServ).Handle)
	api.GET("/buckets/:bucketID/objects", controllers.NewGetObjectsInBucket(getObjectsInBucketServ).Handle)
	api.GET("/buckets", controllers.NewGetBuckets(getBucketsService).Handle)
	api.GET("/buckets/:bucketID/buckets", controllers.NewGetBucketsInBucket(getBucketsInBucketServ).Handle)
	api.GET("/objects/:objectID/download", controllers.NewDownloadObject(urlGenerator, generateObjectPathServ).Handle)
	api.DELETE("/objects/:objectID", controllers.NewDeleteObject(deleteObjectServ).Handle)

	e.Logger.Fatal(e.Start(address))
}

package main

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/domain/entities"
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

	buckets := make([]entities.Bucket, 0)
	buckets = append(buckets, entities.Bucket{
		ID:       "48fded16-34e8-45df-993d-6c0e39ca0308",
		Name:     "test",
		ParentID: nil,
	})

	objects := make([]entities.Object, 0)
	objects = append(objects, entities.Object{
		ID:        "4b3622bc-d5ec-4071-927c-b649611cdb18",
		Name:      "imagen.png",
		Extension: ".png",
		BucketID:  "48fded16-34e8-45df-993d-6c0e39ca0308",
	})

	bucketRepository := repositories.NewRamBucketRepository(buckets)
	objectRepository := repositories.NewRamObjectRepository(objects)

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

	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/storage", storageFolderName)
	e.POST("/buckets", controllers.NewCreateBucket(createBucketServ).Handle)
	e.POST("/buckets/:bucketID/objects", controllers.NewCreateObject(createObjectServ).Handle)
	e.GET("/buckets", controllers.NewGetBuckets(bucketRepository).Handle)
	e.GET("/objects/:objectID/download", controllers.NewDownloadObject(urlGenerator, generateObjectPathServ).Handle)

	e.Logger.Fatal(e.Start(address))
}

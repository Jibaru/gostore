package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetObjectsInBucket struct {
	getObjectsInBucketService application.GetObjectsInBucketServiceInputPort
}

type GetObjectsInBucketParams struct {
	BucketID string `param:"bucketID"`
}

func NewGetObjectsInBucket(
	getObjectsInBucketService application.GetObjectsInBucketServiceInputPort,
) *GetObjectsInBucket {
	return &GetObjectsInBucket{
		getObjectsInBucketService,
	}
}

func (ctrl *GetObjectsInBucket) Handle(c echo.Context) error {
	params := GetObjectsInBucketParams{}
	err := (&echo.DefaultBinder{}).BindPathParams(c, &params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	objects, err := ctrl.getObjectsInBucketService.Do(params.BucketID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, objects)
}

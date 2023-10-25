package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeleteBucket struct {
	deleteBucketService application.DeleteBucketServiceInputPort
}

type DeleteBucketParams struct {
	BucketID string `param:"bucketID"`
}

func NewDeleteBucket(
	deleteBucketService application.DeleteBucketServiceInputPort,
) *DeleteBucket {
	return &DeleteBucket{
		deleteBucketService,
	}
}

func (ctrl *DeleteBucket) Handle(c echo.Context) error {
	params := DeleteBucketParams{}
	err := (&echo.DefaultBinder{}).BindPathParams(c, &params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	err = ctrl.deleteBucketService.Do(params.BucketID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, dtos.NewMessage("bucket deleted"))
}

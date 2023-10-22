package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetBucketsInBucket struct {
	getBucketsInBucketService application.GetBucketsInBucketServiceInputPort
}

type GetBucketsInBucketParams struct {
	BucketID string `param:"bucketID"`
}

func NewGetBucketsInBucket(
	getBucketsInBucketService application.GetBucketsInBucketServiceInputPort,
) *GetBucketsInBucket {
	return &GetBucketsInBucket{
		getBucketsInBucketService,
	}
}

func (ctrl *GetBucketsInBucket) Handle(c echo.Context) error {
	params := GetBucketsInBucketParams{}
	err := (&echo.DefaultBinder{}).BindPathParams(c, &params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	buckets, err := ctrl.getBucketsInBucketService.Do(params.BucketID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, buckets)
}

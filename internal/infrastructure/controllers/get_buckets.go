package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetBuckets struct {
	getBucketsService application.GetBucketServiceInputPort
}

func NewGetBuckets(
	getBucketsService application.GetBucketServiceInputPort,
) *GetBuckets {
	return &GetBuckets{
		getBucketsService,
	}
}

func (ctrl *GetBuckets) Handle(c echo.Context) error {
	buckets, err := ctrl.getBucketsService.Do()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, buckets)
}

package controllers

import (
	"github.com/jibaru/gostore/internal/domain/repositories"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetBuckets struct {
	bucketRepository repositories.BucketRepository
}

func NewGetBuckets(
	bucketRepository repositories.BucketRepository,
) *GetBuckets {
	return &GetBuckets{bucketRepository}
}

func (ctrl *GetBuckets) Handle(c echo.Context) error {
	buckets, err := ctrl.bucketRepository.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, buckets)
}

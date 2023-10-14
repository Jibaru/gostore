package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateBucket struct {
	createBucketService *application.CreateBucketService
}

type CreateBucketRequest struct {
	Name     string  `form:"name" json:"name" xml:"name"`
	ParentID *string `form:"parent_id" json:"parent_id" xml:"parent_id"`
}

func NewCreateBucket(
	createBucketService *application.CreateBucketService,
) *CreateBucket {
	return &CreateBucket{
		createBucketService,
	}
}

func (ctrl *CreateBucket) Handle(c echo.Context) error {
	request := CreateBucketRequest{}

	err := (&echo.DefaultBinder{}).BindBody(c, &request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	bucket, err := ctrl.createBucketService.Do(request.Name, request.ParentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusCreated, bucket)
}

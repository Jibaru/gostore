package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateObject struct {
	createObjectService *application.CreateObjectService
}

type CreateObjectParams struct {
	BucketID string `param:"bucketID"`
}

func NewCreateObject(
	createObjectService *application.CreateObjectService,
) *CreateObject {
	return &CreateObject{
		createObjectService,
	}
}

func (ctrl *CreateObject) Handle(c echo.Context) error {
	params := CreateObjectParams{}
	err := (&echo.DefaultBinder{}).BindPathParams(c, &params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	object, err := ctrl.createObjectService.Do(file, params.BucketID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusCreated, object)
}

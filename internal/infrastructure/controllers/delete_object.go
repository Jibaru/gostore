package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeleteObject struct {
	deleteObjectService application.DeleteObjectServiceInputPort
}

type DeleteObjectParams struct {
	ObjectID string `param:"objectID"`
}

func NewDeleteObject(
	deleteObjectService application.DeleteObjectServiceInputPort,
) *DeleteObject {
	return &DeleteObject{
		deleteObjectService,
	}
}

func (ctrl *DeleteObject) Handle(c echo.Context) error {
	params := DeleteObjectParams{}
	err := (&echo.DefaultBinder{}).BindPathParams(c, &params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	err = ctrl.deleteObjectService.Do(params.ObjectID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, dtos.NewMessage("object deleted"))
}

package controllers

import (
	"github.com/jibaru/gostore/internal/application"
	"github.com/jibaru/gostore/internal/infrastructure/controllers/dtos"
	"github.com/jibaru/gostore/internal/shared"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DownloadObject struct {
	urlGenerator              *shared.UrlGenerator
	generateObjectPathService *application.GenerateObjectPathService
}

type DownloadObjectParams struct {
	ObjectID string `param:"objectID"`
}

func NewDownloadObject(
	urlGenerator *shared.UrlGenerator,
	generateObjectPathService *application.GenerateObjectPathService,
) *DownloadObject {
	return &DownloadObject{
		urlGenerator,
		generateObjectPathService,
	}
}

func (ctrl *DownloadObject) Handle(c echo.Context) error {
	params := DownloadObjectParams{}
	err := (&echo.DefaultBinder{}).BindPathParams(c, &params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	path, err := ctrl.generateObjectPathService.Do(params.ObjectID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewMessage(err.Error()))
	}

	return c.Redirect(http.StatusSeeOther, ctrl.urlGenerator.GenerateUrlFromObjectPath(path))
}

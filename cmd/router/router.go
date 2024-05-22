package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pinares/config"
	"pinares/internal/service"
	"pinares/pkg/utilities"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET(configuration.Config_.PrefixPath+"/__health", healthHandler)

	e.GET(configuration.Config_.PrefixPath+"/tower", service.GetTower)
	e.GET(configuration.Config_.PrefixPath+"/towers", service.GetTowers)
	e.POST(configuration.Config_.PrefixPath+"/tower", service.PostTower)

	utilities.SetLoggerMiddleware(e)
	return e
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

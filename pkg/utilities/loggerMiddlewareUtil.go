package utilities

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func skipper(c echo.Context) bool {
	return strings.Contains(c.Request().RequestURI, "__health")
}

func SetLoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: skipper,
		Format: `{time: ${time_custom}, host: ${host}, method: ${method}, uri: ${uri}, status: ${status}, ` +
			`error: ${error}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
}

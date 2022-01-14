package httpdel

import (
	"github.com/labstack/echo"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser"
)

func MapRoutes(fibGroup *echo.Group, h parser.Handler) {
	fibGroup.GET("/get", h.Get())
}
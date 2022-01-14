package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/delivery/http_del"
)

func (s *Server) MapHandlers(e *echo.Echo, handler parser.Handler) error {

	api := e.Group("/api")

	health := api.Group("/health")
	
	httpdel.MapRoutes(api, handler)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
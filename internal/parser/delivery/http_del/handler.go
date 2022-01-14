package httpdel

import (
	"net/http"

	"github.com/labstack/echo"
)

type ParseHandler struct {
	
}

func NewParseHandler() *ParseHandler {
	return &ParseHandler{}
}

func (h *ParseHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		

		return c.JSON(http.StatusOK, "HI THERE")
	}
}
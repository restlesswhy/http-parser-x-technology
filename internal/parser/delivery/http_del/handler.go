package httpdel

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser"
)

type ParseHandler struct {
	parserUC parser.UseCase
}

func NewParseHandler(parserUC parser.UseCase) *ParseHandler {
	return &ParseHandler{
		parserUC: parserUC,
	}
}

func (h *ParseHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		x, err := h.parserUC.Get(c.Request().Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, x)
	}
}
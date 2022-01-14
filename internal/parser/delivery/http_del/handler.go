package httpdel

import (
	"context"
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
		x, _ := h.parserUC.Get(context.Background())

		return c.JSON(http.StatusOK, x)
	}
}
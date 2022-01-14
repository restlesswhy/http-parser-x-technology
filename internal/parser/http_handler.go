package parser

import "github.com/labstack/echo"

type Handler interface {
	Get() echo.HandlerFunc
}
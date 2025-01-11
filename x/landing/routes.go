package landing

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/hway/x/landing/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	e.HTTPErrorHandler = handlers.ErrorHandler

	e.GET("/", handlers.IndexHandler)
	handlers.RegisterHandler(e.Group("/register"))
}

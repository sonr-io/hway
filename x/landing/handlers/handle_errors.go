package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/hway/pkg/context"
	"github.com/onsonr/hway/x/landing/views"
)

func ErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		if he.Code == 500 {
			// Log the error if needed
			c.Logger().Errorf("Error: %v", he.Message)
			context.Render(c, views.ErrorView(he.Error()))
		}
	}
}

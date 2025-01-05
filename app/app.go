// Package gateway provides the default routes for the Sonr hway.
package app

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/hway/gateway/handlers"
	config "github.com/onsonr/hway/internal/config"
	hwayorm "github.com/onsonr/hway/internal/models"
	"github.com/onsonr/hway/pkg/common"
	"github.com/onsonr/hway/pkg/context"
)

type Gateway = *echo.Echo

// New returns a new Gateway instance
func New(env config.Hway, ipc common.IPFS, dbq *hwayorm.Queries) (Gateway, error) {
	e := echo.New()

	// Built-in middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.Use(context.UseGateway(env, ipc, dbq))

	// Register View Handlers
	e.HTTPErrorHandler = handlers.ErrorHandler
	e.GET("/", handlers.IndexHandler)
	handlers.RegisterHandler(e.Group("/register"))
	return e, nil
}

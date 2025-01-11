package app

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/hway/pkg/common"
	config "github.com/onsonr/hway/pkg/config"
	"github.com/onsonr/hway/pkg/context"
	hwayorm "github.com/onsonr/hway/pkg/models"
	"github.com/onsonr/hway/x/landing"
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

	// Use the gateway middleware
	e.Use(context.UseGateway(env, ipc, dbq))

	// Add Module Routes
	landing.RegisterRoutes(e)
	return e, nil
}

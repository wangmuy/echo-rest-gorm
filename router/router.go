package router

import (
	"echo-rest-gorm/api"
	"echo-rest-gorm/config"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo, c *config.Config) *echo.Echo {
	dbGroup := e.Group("/db")
	api.DbGroup(dbGroup)
	return e
}

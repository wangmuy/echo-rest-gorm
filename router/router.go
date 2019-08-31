package router

import (
	"echo-rest-gorm/api"
	"echo-rest-gorm/config"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo, c *config.Config) *echo.Echo {
	if c.DB_CONN != "" {
		dbGroup := e.Group("/db")
		api.DbGroup(dbGroup)
	}
	printConnGroup := e.Group("/printconn")
	api.PrintConnGroup(printConnGroup)
	return e
}

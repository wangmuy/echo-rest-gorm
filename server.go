package main

import (
	"echo-rest-gorm/config"

	"github.com/labstack/echo"

	"echo-rest-gorm/db"
	"echo-rest-gorm/router"
)

func main() {
	c := config.GetConfig()
	db.Init(c)
	e := echo.New()
	router.Init(e, c)

	e.Logger.Fatal(e.Start(":1323"))
}

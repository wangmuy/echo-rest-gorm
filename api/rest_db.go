package api

import (
	"net/http"

	"github.com/labstack/echo"

	"echo-rest-gorm/db"
	"echo-rest-gorm/model"
)

func DbGroup(g *echo.Group) {
	g.GET("/invoice/:id", GetInvoiceOne)
	g.GET("/employees", GetEmployeeList)
}

func GetInvoiceOne(c echo.Context) error {
	id := c.Param("id")
	db := db.DBManager()
	invoice := model.Invoice{}
	db.Where("InvoiceId=?", id).First(&invoice)
	return c.JSON(http.StatusOK, invoice)
}

func GetEmployeeList(c echo.Context) error {
	db := db.DBManager()
	employees := []model.Employee{}
	db.Find(&employees)
	return c.JSON(http.StatusOK, employees)
}

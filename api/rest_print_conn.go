package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"echo-rest-gorm/model"

	"github.com/labstack/echo"
)

func PrintConnGroup(g *echo.Group) {
	g.GET("/get", PrintGetConn)
	g.POST("/post", PrintNonGetConn)
	g.PUT("/put", PrintNonGetConn)
	g.DELETE("/delete", PrintNonGetConn)
}

func PrintGetConn(c echo.Context) error {
	connInfo := model.ConnectionInfo{}
	request := c.Request()

	headerList := getHeaderList(request)
	connInfo.Header = strings.Join(headerList, ", ")

	paramList := getQueryParamList(c.QueryParams())
	connInfo.Params = strings.Join(paramList, ", ")

	connInfo.Path = c.Path()
	connInfo.Method = request.Method
	json := c.JSONPretty(http.StatusOK, connInfo, "  ")
	fmt.Printf("%+v\n", connInfo)
	return json
}

func PrintNonGetConn(c echo.Context) error {
	connInfo := model.ConnectionInfo{}
	request := c.Request()

	headerList := getHeaderList(request)
	connInfo.Header = strings.Join(headerList, ", ")

	queryMap := echo.Map{}
	c.Bind(&queryMap)
	formValues, _ := c.FormParams()
	formValueList := getQueryParamList(formValues)
	paramList := getBodyParamList(queryMap)
	paramList = append(paramList, formValueList...)
	connInfo.Params = strings.Join(paramList, ", ")

	connInfo.Path = c.Path()
	connInfo.Method = request.Method
	json := c.JSONPretty(http.StatusOK, connInfo, "  ")
	fmt.Printf("%+v\n", connInfo)
	return json
}

func getHeaderList(request *http.Request) []string {
	var headerList []string
	for name, headers := range request.Header {
		for _, h := range headers {
			headerList = append(headerList, fmt.Sprintf("%v=%v", name, h))
		}
	}
	return headerList
}

func getQueryParamList(paramMap url.Values) []string {
	var paramList []string
	for key, valueList := range paramMap {
		for _, v := range valueList {
			paramList = append(paramList, fmt.Sprintf("%v=%v", key, v))
		}
	}
	return paramList
}

func getBodyParamList(paramMap echo.Map) []string {
	var paramList []string
	for k, v := range paramMap {
		paramList = append(paramList, fmt.Sprintf("%v=%v", k, v))
	}
	return paramList
}

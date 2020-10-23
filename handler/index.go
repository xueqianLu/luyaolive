package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetIndex(c echo.Context) error {
	return c.HTML(http.StatusOK, "<p>我是木木，来自北京</p><p>爸爸是一名软件工程师</p>")
}

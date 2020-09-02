package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type GetResonseData struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

func main() {
	e := echo.New()

	data := GetResonseData{
		Status: http.StatusOK,
		Body:   "hello echo",
	}
	res, err := json.Marshal(data)
	if err != nil {
		data.Status = http.StatusInternalServerError
	}

	resMap := map[string]interface{}{
		"status": http.StatusOK,
		"body":   "hello echo",
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, resMap)
	})
	e.GET("/struct", func(c echo.Context) error {
		return c.JSON(http.StatusOK, string(res))
	})

	e.Start(":8080")
}

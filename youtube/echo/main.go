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

func mapResponse(c echo.Context) error {
	resMap := map[string]interface{}{
		"status": http.StatusOK,
		"body":   "hello echo",
	}

	return c.JSON(http.StatusOK, resMap)
}

func structJsonMarshalResponse(c echo.Context) error {
	data := GetResonseData{
		Status: http.StatusOK,
		Body:   "hello echo",
	}
	res, err := json.Marshal(data)
	if err != nil {
		data.Status = http.StatusInternalServerError
	}

	return c.JSON(data.Status, string(res))
}

func main() {
	e := echo.New()

	e.GET("/", mapResponse)
	e.GET("/struct", structJsonMarshalResponse)

	e.Start(":8080")
}

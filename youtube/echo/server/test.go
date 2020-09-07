package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type GetResponseData struct {
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
	data := GetResponseData{
		Status: http.StatusOK,
		Body:   "hello echo",
	}
	// どっちでもおｋ
	// bytes, err := json.Marshal(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		data.Status = http.StatusInternalServerError
	}
	log.Printf("status: %d", data.Status)

	return c.JSONBlob(data.Status, bytes)
}

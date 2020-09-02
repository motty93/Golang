package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type GetResponseData struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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

	return c.JSON(data.Status, string(bytes))
}

func productsResponse(c echo.Context) error {
	data := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}
	var product Product
	for _, p := range data {
		for key := range p {
			pID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			if pID == key {
				product = Product{Id: key, Name: p[key]}
			}
		}
	}
	if product.Id == 0 {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	bytes, _ := json.Marshal(product)

	return c.JSON(http.StatusOK, string(bytes))
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo framework")
	})
	e.GET("/map", mapResponse)
	e.GET("/struct", structJsonMarshalResponse)
	e.GET("/products/:id", productsResponse)

	e.Logger.Print("Listening on port 8080...")
	e.Logger.Fatal(e.Start(":8080"))
}

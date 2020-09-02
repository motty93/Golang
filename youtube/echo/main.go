package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

var data = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

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

func productsIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, data)
}

func productShow(c echo.Context) error {
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
	bytes, _ := json.Marshal(product)
	if product.Id == 0 {
		// return c.JSONBlob(http.StatusNotFound, bytes)
		return c.JSON(http.StatusNotFound, map[string]string{"message": "product not found"})
	}

	return c.JSONBlob(http.StatusOK, bytes)
}

func productCreate(c echo.Context) error {
	var reqBody Product
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	product := map[int]string{
		len(data) + 1: reqBody.Name,
	}
	data = append(data, product)
	return c.JSON(http.StatusOK, product)
}

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo framework")
	})
	e.GET("/map", mapResponse)
	e.GET("/struct", structJsonMarshalResponse)
	e.GET("/products", productsIndex)
	e.GET("/product/:id", productShow)
	e.POST("/product", productCreate)

	e.Logger.Printf("Listening on port %s...", port)
	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}

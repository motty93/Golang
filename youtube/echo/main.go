package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type GetResponseData struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

// Product echo REST test struct
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=4"`
	// Vendor          string `json:"vendor" validate:"min=5,max=10"`
	// Email           string `json:"email" validate:"required_with=Vendor,email"`
	// Website         string `json:"website" validate:"url"`
	// Country         string `json:"country" validate:"len=2"`
	// DefaultDeviceIP string `json:"default_device_ip" validate:"ip"`
}

// ProductValidator echo validator for product
type ProductValidator struct {
	validator *validator.Validate
}

// Validate validates product request body
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var (
	e    *echo.Echo
	v    *validator.Validate
	data []map[int]string
	port string
)

func init() {
	e = echo.New()
	v = validator.New()
	port = os.Getenv("MY_APP_PORT")
	data = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}
	if port == "" {
		port = "8080"
	}
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
				product = Product{ID: key, Name: p[key]}
			}
		}
	}
	bytes, _ := json.Marshal(product)
	if product.ID == 0 {
		// return c.JSONBlob(http.StatusNotFound, bytes)
		return c.JSON(http.StatusNotFound, map[string]string{"message": "product not found"})
	}

	return c.JSONBlob(http.StatusOK, bytes)
}

func productCreate(c echo.Context) error {
	var reqBody Product
	e.Validator = &ProductValidator{validator: v}

	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	// add params validator
	// if err := v.Struct(reqBody); err != nil {
	// 	return err
	// }
	// echo validator
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	product := map[int]string{
		len(data) + 1: reqBody.Name,
	}
	data = append(data, product)
	return c.JSON(http.StatusOK, product)
}

func productUpdate(c echo.Context) error {
	var product Product
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for _, p := range data {
		for key := range p {
			if pID == key {
				product = Product{ID: key, Name: p[key]}
			}
		}
	}

	if product.ID == 0 {
		// return c.JSONBlob(http.StatusNotFound, bytes)
		return c.JSON(http.StatusNotFound, map[string]string{"message": "product not found"})
	}

	var reqBody Product
	e.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	product.Name = reqBody.Name
	bytes, _ := json.Marshal(product)
	return c.JSONBlob(http.StatusOK, bytes)
}

func main() {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo framework")
	})
	e.GET("/map", mapResponse)
	e.GET("/struct", structJsonMarshalResponse)
	e.GET("/products", productsIndex)
	e.GET("/products/:id", productShow)
	e.POST("/product", productCreate)
	e.PUT("/products/:id", productUpdate)

	e.GET("/tests", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []string{"mobile", "tv", "oven"})
	})
	e.GET("/tests/:vendor", func(c echo.Context) error {
		// return c.JSON(http.StatusOK, c.Param("vendor"))
		// query parameterを取得する
		return c.JSON(http.StatusOK, c.QueryParam("olderThan"))
	})

	e.Logger.Printf("Listening on port %s...", port)
	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}

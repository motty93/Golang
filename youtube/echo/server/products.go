package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

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

var products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

func indexProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func showProduct(c echo.Context) error {
	var product Product
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for _, p := range products {
		for key := range p {
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

func createProduct(c echo.Context) error {
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
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	var product Product
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for _, p := range products {
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

func deleteProduct(c echo.Context) error {
	var product Product
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for i, p := range products {
		for key := range p {
			if pID == key {
				product = Product{ID: key, Name: p[key]}
				index = i
			}
		}
	}

	bytes, _ := json.Marshal(product)
	if product.ID == 0 {
		// return c.JSONBlob(http.StatusNotFound, bytes)
		return c.JSON(http.StatusNotFound, map[string]string{"message": "product not found"})
	}
	// 該当するidの削除処理
	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
		// [1,2,3,4,5,6]
		// [1,2] + [4,5,6] = [1,2,4,5,6]
	}
	products = splice(products, index)

	return c.JSONBlob(http.StatusOK, bytes)
}

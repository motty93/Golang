package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

// Start is routing function
func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo framework")
	})
	e.GET("/map", mapResponse)
	e.GET("/struct", structJsonMarshalResponse)
	e.GET("/products", indexProduct)
	e.GET("/products/:id", showProduct)
	e.POST("/product", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

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

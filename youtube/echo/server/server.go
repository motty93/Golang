package server

import (
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

// Start is routing function
func Start() {

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

	e.Logger.Printf("Listening on port %s...", cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port))) // e.Logger.Fatal(e.Start(":8080"))
}

package server

import (
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

// type MiddlewareFunc func(HandlerFunc) HandlerFunc
// type HandlerFunc func(Context) error

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside serverheader middleware")
		c.Request().URL.Path = "/krunal"
		fmt.Printf("%+v\n", c.Request())
		return next(c)
	}
}

func serverMessageDo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware Do")
		return next(c)
	}
}

// Start starts the application
func Start() {
	// e.Use(serverMessage)
	// e.Pre(serverMessage)

	// スラッシュを削除
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo framework")
	})
	e.GET("/map", mapResponse)
	e.GET("/struct", structJsonMarshalResponse)
	e.GET("/products", indexProduct, serverMessageDo) // 複数可能
	e.GET("/products/:id", showProduct)
	e.POST("/product", createProduct, middleware.BodyLimit("2M")) // 2MまでしかBodyは受け付けない
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

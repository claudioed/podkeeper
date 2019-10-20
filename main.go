package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main()  {

	// Creating echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	data :=make(map[string]string)

	data["status"] = "ok"

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200,data)
	})
	
	// Server
	e.Logger.Fatal(e.Start(":8888"))

}

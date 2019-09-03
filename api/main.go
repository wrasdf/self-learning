package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)
  e.GET("/employees", EController.all())
  e.GET("/employees/{id}", EController.get(id))
  e.POST("/employees/{id}", EController.add())
  e.PUT("/employees/{id}", EController.update())
  e.DELETE("/employees/{id}", EController.delete(id))

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

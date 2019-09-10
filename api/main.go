package main

import (
  "net/http"
  // "services/employees"

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
  e.GET("/", helloworld)
  e.GET("/health", health)

  // e.GET("/employees", Employees.all())
  // e.GET("/employees/{id}", Employees.get(id))
  // e.POST("/employees/{id}", Employees.add())
  // e.PATCH("/employees/{id}", Employees.patch())
  // e.PUT("/employees/{id}", Employees.update())
  // e.DELETE("/employees/{id}", Employees.delete(id))

  // e.GET("/departments", Departments.all())
  // e.GET("/departments/{id}", Departments.get(id))
  // e.POST("/departments/{id}", Departments.add(id))
  // e.PATCH("/departments/{id}", Departments.patch(id))
  // e.PUT("/departments/{id}", Departments.update(id))
  // e.DELETE("/departments/{id}", Departments.delete(id))

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func helloworld(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func health(c echo.Context) error {
  return c.String(http.StatusOK, "ok!")
}

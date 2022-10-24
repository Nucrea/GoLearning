package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// "gotest/internal/repository"
// "gotest/internal/service"

func main() {
	// userRepo := repository.NewUserRepository()
	// userService := service.NewUserService(userRepo)

	// userService.CreateUser(nil, &{})

	e := echo.New()
	e.GET("/", test)
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
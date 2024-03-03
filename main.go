package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// In production, the secret key of the CookieStore
// and database name would be obtained from a .env file

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("ENV")

	e := echo.New()

	e.Static("/", "assets")

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>This website is under construction</h1>")
	})

	var ports string
	if env == "PRODUCTION" {
		ports = fmt.Sprintf(":%d", 8080)
	} else {
		ports = fmt.Sprintf(":%d", 3002)
	}

	// Start Server
	e.Logger.Fatal(e.Start(ports))

}

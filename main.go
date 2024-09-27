package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, getStringWithTimestamp())
	})
	e.Logger.Fatal(e.Start(":8080"))
}

// getStringWithTimestamp returns the string with "OK" and the current timestamp
func getStringWithTimestamp() string {
	return "OK - " + time.Now().Format(time.RFC3339)
}

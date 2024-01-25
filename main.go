package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		min, _ := strconv.Atoi(c.QueryParam("min"))
		max, _ := strconv.Atoi(c.QueryParam("max"))

		if min >= max {
			return c.String(http.StatusBadRequest, "Invalid range!")
		}

		num := rand.Intn(max-min) + min

		return c.String(http.StatusOK, fmt.Sprintf("Random number between %d and %d: %d", min, max, num))
	})

	e.Logger.Fatal(e.Start(":8080"))
}

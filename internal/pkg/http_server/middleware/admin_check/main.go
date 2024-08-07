package admin_check

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func New(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("User-Role") == "admin" {
			fmt.Println("red button user detected")
		}
		return next(c)
	}
}

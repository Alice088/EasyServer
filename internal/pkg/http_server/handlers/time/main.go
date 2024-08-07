package time

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func New(c echo.Context) error {
	t, err := time.Parse("2006-Jan-02", "2025-Jan-01")
	if err != nil {
		c.Logger().Errorf("internal.pkg.http_server.handlers.time.main.New: " + err.Error())
	}

	timeUntil := t.Sub(time.Now()).Hours() / 24

	return c.String(http.StatusOK, strconv.Itoa(int(timeUntil)))
}

package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// GetTime 시간을가져오는 메소드
func GetTime(c echo.Context) error {
	return c.String(http.StatusOK, strconv.FormatInt(time.Now().Unix(), 10))
}

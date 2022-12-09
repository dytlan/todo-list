package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewSuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

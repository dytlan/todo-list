package functions

import (
	"github.com/dytlan/moonlay-todo-list/engine/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func BindAndValidate(c echo.Context, req interface{}) error {
	if err := c.Bind(&req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return nil
}

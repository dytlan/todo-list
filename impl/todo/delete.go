package todo

import (
	"github.com/dytlan/moonlay-todo-list/engine/response"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/impl/functions"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ths *service) Delete(c echo.Context) error {
	var deleteReq DeleteReq
	if err := functions.BindAndValidate(c, &deleteReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, err := ths.accessor.Delete(c, accessor.DeleteSpec{
		IDs: deleteReq.IDs,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, DeleteRes{
		AffectedRows: res.AffectedRows,
	})
}

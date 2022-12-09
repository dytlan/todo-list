package todo

import (
	"github.com/dytlan/moonlay-todo-list/engine/response"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/impl/functions"
	"github.com/labstack/echo/v4"
)

func (ths *service) List(c echo.Context) error {
	var listReq ListReq
	if err := functions.BindAndValidate(c, &listReq); err != nil {
		return err
	}
	res, err := ths.accessor.List(c, accessor.ListSpec{
		Filter:   listReq.Filter,
		PageSpec: listReq.PageSpec,
	})
	if err != nil {
		return err
	}
	return response.NewSuccessResponse(c, ListRes{
		Data:     res.Data,
		PageInfo: res.PageInfo,
	})
}

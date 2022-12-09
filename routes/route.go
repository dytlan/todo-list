package routes

import (
	"github.com/dytlan/moonlay-todo-list/impl/todo"
	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo, svc todo.Service) {
	e.GET("/list", svc.List)
	e.POST("/upsert", svc.Upsert)
	e.DELETE("/delete", svc.Delete)
	e.GET("/getByIds", svc.GetByIDs)
}

package todo

import (
	"github.com/dytlan/moonlay-todo-list/engine/file"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/labstack/echo/v4"
)

type Service interface {
	List(c echo.Context) error
	Upsert(c echo.Context) error
	Delete(c echo.Context) error
	GetByIDs(c echo.Context) error
}

type service struct {
	accessor accessor.Accessor
	file     file.Service
}

func NewService(accessor accessor.Accessor, file file.Service) Service {
	return &service{
		accessor: accessor,
		file:     file,
	}
}

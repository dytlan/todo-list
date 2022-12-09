package todo

import (
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/impl/functions"
)

type ListRes struct {
	Data     []accessor.ToDoDB  `json:"data"`
	PageInfo functions.PageInfo `json:"pageInfo"`
}

type ListReq struct {
	Filter   accessor.FilterCriteriaSpec `json:"filter"`
	PageSpec functions.PageSpec          `json:"pageSpec"`
}

type UpsertReq struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" validate:"required,lte=100,alphanum"`
	Description string `json:"description" validate:"required,lte=1000"`
	File        string `json:"file" validate:"base64"`
	Level       int    `json:"level" validate:"required,gte=1,lte=3"`
	IsDone      *bool  `json:"isDone"`
	MasterID    uint   `json:"masterId"`
	RootID      uint   `json:"rootId"`
}

type UpsertRes struct {
	AffectedRows int64 `json:"affectedRows"`
}

type DeleteReq struct {
	IDs []uint `json:"Ids" validate:"required"`
}

type DeleteRes struct {
	AffectedRows int64 `json:"affectedRows"`
}

type GetByIDsReq struct {
	IDs    []uint                  `json:"Ids"`
	Filter accessor.GetByIDsFilter `json:"filter"`
}

type GetByIDsRes struct {
	Data []GetByIDsData `json:"data"`
}

type GetByIDsData struct {
	accessor.ToDoDB
	SubData []*GetByIDsData `json:"subData"`
}

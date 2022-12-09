package accessor

import (
	"github.com/dytlan/moonlay-todo-list/engine/response"
	"github.com/dytlan/moonlay-todo-list/impl/functions"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type ToDoAccessor interface {
	List(c echo.Context, listSpec ListSpec) (ListRes, error)
	Upsert(c echo.Context, spec UpsertSpec) (UpsertRes, error)
	GetByID(c echo.Context, spec GetByIDSpec) (ToDoDB, error)
	Delete(c echo.Context, spec DeleteSpec) (DeleteRes, error)
	GetByIDs(c echo.Context, spec GetByIDsSpec) ([]ToDoDB, error)
}

type ToDoDB struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
	Level       int    `json:"level"`
	IsDone      *bool  `json:"isDone"`
	MasterID    uint   `json:"masterId"`
	RootID      uint   `json:"rootId"`
}

type FilterCriteriaSpec struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Level     int    `json:"level"`
	IsDone    *bool  `json:"isDone"`
}

type ListSpec struct {
	Filter   FilterCriteriaSpec
	PageSpec functions.PageSpec
}

type ListRes struct {
	Data     []ToDoDB
	PageInfo functions.PageInfo
}

type UpsertSpec struct {
	ID          uint
	Title       string
	Description string
	File        string
	Level       int
	IsDone      *bool
	MasterID    uint
	RootID      uint
}

type UpsertRes struct {
	AffectedRow int64
}

type GetByIDSpec struct {
	ID uint
}

type DeleteSpec struct {
	IDs []uint
}

type DeleteRes struct {
	AffectedRows int64
}

type GetByIDsSpec struct {
	IDs    []uint
	Filter GetByIDsFilter
}

type GetByIDsFilter struct {
	FilterCriteriaSpec
	ShowTreeData bool `json:"showTreeData"`
}

type GetByIDsRes struct {
	Data []ToDoDB
}

func (ths *accessor) filterQuery(spec FilterCriteriaSpec, query *gorm.DB) *gorm.DB {
	if spec.StartDate != "" && spec.EndDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", spec.StartDate, spec.EndDate)
	}
	if spec.IsDone != nil {
		query = query.Where("is_done = ?", *spec.IsDone)
	}
	if spec.Level != 0 {
		query = query.Where("level = ?", spec.Level)
	}
	return query
}

func (ths *accessor) List(c echo.Context, spec ListSpec) (res ListRes, err error) {
	query := ths.db.Model(&ToDoDB{}).Order("created_at DESC")
	if spec.Filter != (FilterCriteriaSpec{}) {
		query = ths.filterQuery(spec.Filter, query)
	}
	if spec.PageSpec != (functions.PageSpec{}) {
		query = query.Limit(spec.PageSpec.LimitValue()).Offset(spec.PageSpec.OffsetValue())
	}

	if err = query.Count(&res.PageInfo.TotalItems).Error; err != nil {
		return ListRes{}, response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
			Key:   "spec",
			Value: spec,
		})
	}

	if err = query.Find(&res.Data).Error; err != nil {
		return ListRes{}, response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
			Key:   "spec",
			Value: spec,
		})
	}

	res.PageInfo = spec.PageSpec.PageInfo(res.PageInfo.TotalItems)
	return
}

func (ths *accessor) Upsert(c echo.Context, spec UpsertSpec) (res UpsertRes, err error) {
	query := ths.db.Model(&ToDoDB{})
	if spec.ID != 0 {
		query = query.Where("id = ?", spec.ID).Select("title", "description", "file", "level", "is_done", "master_id").Updates(ToDoDB{
			Title:       spec.Title,
			Description: spec.Description,
			File:        spec.File,
			Level:       spec.Level,
			IsDone:      spec.IsDone,
			MasterID:    spec.MasterID,
			RootID:      spec.RootID,
		})
	} else {
		query = query.Create(&ToDoDB{
			Title:       spec.Title,
			Description: spec.Description,
			File:        spec.File,
			Level:       spec.Level,
			IsDone:      spec.IsDone,
			MasterID:    spec.MasterID,
			RootID:      spec.RootID,
		})
	}
	if err = query.Error; err != nil {
		return UpsertRes{}, response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
			Key:   "spec",
			Value: spec,
		})
	}
	res.AffectedRow = query.RowsAffected
	return
}

func (ths *accessor) GetByID(c echo.Context, spec GetByIDSpec) (res ToDoDB, err error) {
	if err = ths.db.Where("id = ?", spec.ID).First(&res).Error; err != nil {
		return res, response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
			Key:   "spec",
			Value: spec,
		})
	}
	return
}

func (ths *accessor) Delete(c echo.Context, spec DeleteSpec) (res DeleteRes, err error) {
	var data []ToDoDB
	query := ths.db.Delete(&data, spec.IDs)
	res.AffectedRows = query.RowsAffected
	if err = query.Error; err != nil {
		return res, response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
			Key:   "spec",
			Value: spec,
		})
	}
	return
}

func (ths *accessor) GetByIDs(c echo.Context, spec GetByIDsSpec) (res []ToDoDB, err error) {
	query := ths.db.Order("level").Where("id IN ?", spec.IDs)
	if spec.Filter != (GetByIDsFilter{}) {
		query = ths.filterQuery(spec.Filter.FilterCriteriaSpec, query)
	}
	if spec.Filter.ShowTreeData {
		query = query.Or("root_id IN ?", spec.IDs)
	}
	if err = query.Order("id DESC").Find(&res).Error; err != nil {
		return res, response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
			Key:   "spec",
			Value: spec,
		})
	}
	return
}

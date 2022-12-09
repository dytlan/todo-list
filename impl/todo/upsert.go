package todo

import (
	"github.com/dytlan/moonlay-todo-list/engine/file"
	"github.com/dytlan/moonlay-todo-list/engine/response"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/impl/functions"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validMimeType = []string{"text/plain", "application/pdf"}

func (ths *service) Upsert(c echo.Context) error {
	var upsertReq UpsertReq
	if err := functions.BindAndValidate(c, &upsertReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	upsertSpec := accessor.UpsertSpec{
		ID:          upsertReq.ID,
		Title:       upsertReq.Title,
		Description: upsertReq.Description,
		Level:       upsertReq.Level,
		IsDone:      upsertReq.IsDone,
		MasterID:    upsertReq.MasterID,
		RootID:      upsertReq.RootID,
	}
	if upsertReq.File != "" {
		bufferFile, fileInformation, err := file.ConstructBase64ToBuffer(upsertReq.File)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err = file.ValidateMimeType(fileInformation.MIMEType, validMimeType); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if upsertSpec.ID != 0 {
			todo, err := ths.accessor.GetByID(c, accessor.GetByIDSpec{
				ID: upsertReq.ID,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			if err = ths.file.DeleteFile(todo.File); err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
		}
		upsertSpec.File, err = ths.file.Base64Upload(bufferFile, fileInformation)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	res, err := ths.accessor.Upsert(c, upsertSpec)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, UpsertRes{
		AffectedRows: res.AffectedRow,
	})
}

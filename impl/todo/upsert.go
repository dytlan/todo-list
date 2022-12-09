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
		return err
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
			return response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
				Key:   "spec",
				Value: upsertReq,
			})
		}
		if err = file.ValidateMimeType(fileInformation.MIMEType, validMimeType); err != nil {
			return response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
				Key:   "spec",
				Value: upsertReq,
			})
		}
		if upsertSpec.ID != 0 {
			todo, err := ths.accessor.GetByID(c, accessor.GetByIDSpec{
				ID: upsertReq.ID,
			})
			if err != nil {
				return err
			}
			if err = ths.file.DeleteFile(todo.File); err != nil {
				return err
			}
		}
		upsertSpec.File, err = ths.file.Base64Upload(bufferFile, fileInformation)
		if err != nil {
			return response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
				Key:   "spec",
				Value: upsertReq,
			})
		}
	}

	res, err := ths.accessor.Upsert(c, upsertSpec)
	if err != nil {
		return response.NewErrorResponseWithMessageAndData(c, http.StatusBadRequest, err.Error(), response.ErrorData{
			Key:   "spec",
			Value: upsertReq,
		})
	}

	return response.NewSuccessResponse(c, UpsertRes{
		AffectedRows: res.AffectedRow,
	})
}

package response

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type ErrorData struct {
	Key   string
	Value interface{}
}

func NewErrorResponse(c echo.Context, statusCode int, err string) error {
	return NewErrorResponseWithMessageAndData(c, statusCode, err)
}

func NewErrorResponseWithMessageAndData(c echo.Context, statusCode int, err string, errorData ...ErrorData) error {
	data := make(map[string]interface{})
	for _, d := range errorData {
		data[d.Key] = d.Value
	}
	return c.JSON(statusCode, ErrorResponse{
		StatusCode: statusCode,
		Message:    err,
		Data:       data,
	})
}

func (ths ErrorResponse) Error() string {
	return fmt.Sprintf("[%v] %v", ths.StatusCode, ths.Message)
}

func (ths ErrorResponse) Is(target error) bool {
	_, ok := target.(ErrorResponse)
	if !ok {
		return false
	}

	return true
}

func (ths ErrorResponse) As(target interface{}) bool {
	t, ok := target.(*ErrorResponse)
	if !ok {
		return false
	}

	t.StatusCode = ths.StatusCode
	t.Data = ths.Data
	t.Message = ths.Message

	return true
}

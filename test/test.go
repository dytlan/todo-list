package test

import (
	"encoding/json"
	"github.com/dytlan/moonlay-todo-list/engine/validation"
	mockFile "github.com/dytlan/moonlay-todo-list/test/mock/engine/file"
	mockAccessor "github.com/dytlan/moonlay-todo-list/test/mock/impl/accessor"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"net/http/httptest"
	"strings"
)

type AllMockServiceDependencies struct {
	Accessor *mockAccessor.MockAccessor
	File     *mockFile.MockService
}

func GetAllMockServiceDependencies(mockCtrl *gomock.Controller) *AllMockServiceDependencies {
	mockAccessorService := mockAccessor.NewMockAccessor(mockCtrl)
	mockFileService := mockFile.NewMockService(mockCtrl)

	return &AllMockServiceDependencies{
		Accessor: mockAccessorService,
		File:     mockFileService,
	}
}

type MockScenario func(allMock *AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder)

type TestCase struct {
	Scenario       string
	ExpectedResult interface{}
	Mock           MockScenario
}

func NewMockHttpRequest(httpMethod, url, payload string) (c echo.Context, rec *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &validation.Validator{Validator: validator.New()}
	req := httptest.NewRequest(httpMethod, url, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	return
}

func UnmarshallResponse(data []byte, response interface{}) {
	json.Unmarshal(data, response)
}

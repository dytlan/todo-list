package functions

import (
	"encoding/json"
	"github.com/dytlan/moonlay-todo-list/engine/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestCaseFunction struct {
	Scenario       string
	ExpectedResult interface{}
	Mock           MockScenarioFunction
	WantErr        bool
}

type MockScenarioFunction func() (c echo.Context, rec *httptest.ResponseRecorder)

type dataTesting struct {
	ID    uint   `json:"id"`
	Title string `json:"title" validate:"required,lte=100,alphanum"`
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

func TestBindAndValidate(t *testing.T) {
	Convey("Given BindAndValidate instances", t, func() {
		httpMethod := http.MethodPost
		url := "/test123"

		testCases := []TestCaseFunction{
			{
				Scenario:       "errorBinding",
				ExpectedResult: dataTesting{},
				Mock: func() (c echo.Context, rec *httptest.ResponseRecorder) {
					return NewMockHttpRequest(httpMethod, url, `{"id":1,"title":0}`)
				},
				WantErr: true,
			},
			{
				Scenario:       "errorValidate",
				ExpectedResult: dataTesting{},
				Mock: func() (c echo.Context, rec *httptest.ResponseRecorder) {
					return NewMockHttpRequest(httpMethod, url, `{"id":1,"title":"this is fail"}`)
				},
				WantErr: true,
			},
			{
				Scenario: "success",
				ExpectedResult: dataTesting{
					ID:    1,
					Title: "success",
				},
				Mock: func() (c echo.Context, rec *httptest.ResponseRecorder) {
					return NewMockHttpRequest(httpMethod, url, `{"id":1,"title":"success"}`)
				},
				WantErr: false,
			},
		}
		for _, testCase := range testCases {
			t.Run(testCase.Scenario, func(t *testing.T) {
				Convey(testCase.Scenario, t, func() {
					var req dataTesting
					c, rec := testCase.Mock()
					err := BindAndValidate(c, &req)
					json.Unmarshal(rec.Body.Bytes(), &req)
					if testCase.WantErr {
						So(err, ShouldNotBeNil)
					} else {
						So(req, ShouldResemble, testCase.ExpectedResult)
						So(err, ShouldBeNil)
					}
				})
			})
		}
	})
}

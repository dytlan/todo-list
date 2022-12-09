package todo

import (
	"fmt"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/impl/functions"
	"github.com/dytlan/moonlay-todo-list/test"
	"github.com/dytlan/moonlay-todo-list/test/fixtures"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	Convey("Given list instances", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		allMock := test.GetAllMockServiceDependencies(mockCtrl)
		svc := NewService(allMock.Accessor, allMock.File)
		httpMethod := http.MethodGet
		url := "/list"

		testCases := []test.TestCase{
			{
				Scenario:       "errorBindingOrValidate",
				ExpectedResult: ListRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					return test.NewMockHttpRequest(httpMethod, url, fixtures.FailBindListRequest)
				},
			},
			{
				Scenario:       "errorAccessorList",
				ExpectedResult: ListRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().List(gomock.Any(), gomock.Any()).Return(accessor.ListRes{}, fmt.Errorf("something went wrong"))
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessListRequest)
				},
			},
			{
				Scenario: "success",
				ExpectedResult: ListRes{
					Data: []accessor.ToDoDB{fixtures.DummyAccessorToDo},
					PageInfo: functions.PageInfo{
						TotalItems:  1,
						ItemPerPage: 1,
						PageNumber:  1,
					},
				},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().List(gomock.Any(), gomock.Any()).Return(accessor.ListRes{
						Data: []accessor.ToDoDB{
							fixtures.DummyAccessorToDo,
						},
						PageInfo: functions.PageInfo{
							TotalItems:  1,
							ItemPerPage: 1,
							PageNumber:  1,
						},
					}, nil)
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessListRequest)
				},
			},
		}
		for _, testCase := range testCases {
			t.Run(testCase.Scenario, func(t *testing.T) {
				Convey(testCase.Scenario, t, func() {
					var res ListRes
					c, rec := testCase.Mock(allMock)
					_ = svc.List(c)
					test.UnmarshallResponse(rec.Body.Bytes(), &res)
					So(res, ShouldResemble, testCase.ExpectedResult)
				})
			})
		}
	})
}

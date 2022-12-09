package todo

import (
	"fmt"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/test"
	"github.com/dytlan/moonlay-todo-list/test/fixtures"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetByIDs(t *testing.T) {
	Convey("Given getByIDs instances", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		allMock := test.GetAllMockServiceDependencies(mockCtrl)
		svc := NewService(allMock.Accessor, allMock.File)
		httpMethod := http.MethodGet
		url := "/getByIds"

		testCases := []test.TestCase{
			{
				Scenario:       "errorBindingOrValidate",
				ExpectedResult: GetByIDsRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					return test.NewMockHttpRequest(httpMethod, url, fixtures.FailBindGetByIDsRequest)
				},
			},
			{
				Scenario:       "errorAccessorGetByIDs",
				ExpectedResult: GetByIDsRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().GetByIDs(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("something went wrong"))
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessGetByIDsRequest)
				},
			},
			{
				Scenario: "success",
				ExpectedResult: GetByIDsRes{
					Data: []*GetByIDsData{
						{
							ToDoDB:  fixtures.DummyAccessorToDo,
							SubData: nil,
						},
					},
				},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().GetByIDs(gomock.Any(), gomock.Any()).Return([]accessor.ToDoDB{fixtures.DummyAccessorToDo}, nil)
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessGetByIDsRequest)
				},
			},
		}
		for _, testCase := range testCases {
			t.Run(testCase.Scenario, func(t *testing.T) {
				Convey(testCase.Scenario, t, func() {
					var res GetByIDsRes
					c, rec := testCase.Mock(allMock)
					_ = svc.GetByIDs(c)
					test.UnmarshallResponse(rec.Body.Bytes(), &res)
					So(res, ShouldResemble, testCase.ExpectedResult)
				})
			})
		}
	})
}

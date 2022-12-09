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

func TestDelete(t *testing.T) {
	Convey("Given delete instances", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		allMock := test.GetAllMockServiceDependencies(mockCtrl)
		svc := NewService(allMock.Accessor, allMock.File)
		httpMethod := http.MethodDelete
		url := "/delete"

		testCases := []test.TestCase{
			{
				Scenario:       "errorBindingOrValidate",
				ExpectedResult: DeleteRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					return test.NewMockHttpRequest(httpMethod, url, fixtures.FailBindDeleteRequestSpec)
				},
			},
			{
				Scenario:       "errorAccessorDelete",
				ExpectedResult: DeleteRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(accessor.DeleteRes{}, fmt.Errorf("something went wrong"))
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessDeleteRequestSpec)
				},
			},
			{
				Scenario: "success",
				ExpectedResult: DeleteRes{
					AffectedRows: 1,
				},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(accessor.DeleteRes{
						AffectedRows: 1,
					}, nil)
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessDeleteRequestSpec)
				},
			},
		}
		for _, testCase := range testCases {
			t.Run(testCase.Scenario, func(t *testing.T) {
				Convey(testCase.Scenario, t, func() {
					var res DeleteRes
					c, rec := testCase.Mock(allMock)
					_ = svc.Delete(c)
					test.UnmarshallResponse(rec.Body.Bytes(), &res)
					So(res, ShouldResemble, testCase.ExpectedResult)
				})
			})
		}
	})
}

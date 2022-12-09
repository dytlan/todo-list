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

func TestUpsert(t *testing.T) {
	Convey("Given upsert instances", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		allMock := test.GetAllMockServiceDependencies(mockCtrl)
		svc := NewService(allMock.Accessor, allMock.File)
		httpMethod := http.MethodPost
		url := "/upsert"

		testCases := []test.TestCase{
			{
				Scenario:       "errorBindingOrValidate",
				ExpectedResult: UpsertRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					return test.NewMockHttpRequest(httpMethod, url, fixtures.FailBindUpsertRequestSpec)
				},
			},
			{
				Scenario:       "errorValidateMimeTypeFile",
				ExpectedResult: UpsertRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					return test.NewMockHttpRequest(httpMethod, url, fixtures.FailValidateBase64FileUpsertRequestSpec)
				},
			},
			{
				Scenario:       "failAccessorGetByID",
				ExpectedResult: UpsertRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(accessor.ToDoDB{}, fmt.Errorf("something went wrong"))
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessUpsertRequestSpec)
				},
			},
			{
				Scenario:       "failDeleteFile",
				ExpectedResult: UpsertRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(fixtures.DummyAccessorToDo, nil)
					allMock.File.EXPECT().DeleteFile(gomock.Any()).Return(fmt.Errorf("something went wrong"))
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessUpsertRequestSpec)
				},
			},
			{
				Scenario:       "failUploadFile",
				ExpectedResult: UpsertRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(fixtures.DummyAccessorToDo, nil)
					allMock.File.EXPECT().DeleteFile(gomock.Any()).Return(nil)
					allMock.File.EXPECT().Base64Upload(gomock.Any(), gomock.Any()).Return("", fmt.Errorf("something went wrong"))
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessUpsertRequestSpec)
				},
			},
			{
				Scenario:       "failAccessorUpsert",
				ExpectedResult: UpsertRes{},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(fixtures.DummyAccessorToDo, nil)
					allMock.File.EXPECT().DeleteFile(gomock.Any()).Return(nil)
					allMock.File.EXPECT().Base64Upload(gomock.Any(), gomock.Any()).Return("path", nil)
					allMock.Accessor.EXPECT().Upsert(gomock.Any(), gomock.Any()).Return(accessor.UpsertRes{}, fmt.Errorf("something went wrong"))
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessUpsertRequestSpec)
				},
			},
			{
				Scenario: "success",
				ExpectedResult: UpsertRes{
					AffectedRows: 1,
				},
				Mock: func(allMock *test.AllMockServiceDependencies) (c echo.Context, rec *httptest.ResponseRecorder) {
					allMock.Accessor.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(fixtures.DummyAccessorToDo, nil)
					allMock.File.EXPECT().DeleteFile(gomock.Any()).Return(nil)
					allMock.File.EXPECT().Base64Upload(gomock.Any(), gomock.Any()).Return("path", nil)
					allMock.Accessor.EXPECT().Upsert(gomock.Any(), gomock.Any()).Return(accessor.UpsertRes{
						AffectedRow: 1,
					}, nil)
					return test.NewMockHttpRequest(httpMethod, url, fixtures.SuccessUpsertRequestSpec)
				},
			},
		}
		for _, testCase := range testCases {
			t.Run(testCase.Scenario, func(t *testing.T) {
				Convey(testCase.Scenario, t, func() {
					var res UpsertRes
					c, rec := testCase.Mock(allMock)
					_ = svc.Upsert(c)
					test.UnmarshallResponse(rec.Body.Bytes(), &res)
					So(res, ShouldResemble, testCase.ExpectedResult)
				})
			})
		}
	})
}

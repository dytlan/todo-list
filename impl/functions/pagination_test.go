package functions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPagination(t *testing.T) {
	Convey("Given BindAndValidate instances", t, func() {
		pagination := PageSpec{
			PageNumber:  0,
			ItemPerPage: 0,
		}
		Convey("return offset value", func() {
			offset := pagination.OffsetValue()
			So(offset, ShouldEqual, 0)
		})
		Convey("return limit value", func() {
			offset := pagination.LimitValue()
			So(offset, ShouldEqual, PaginationDefaultItemPerPage)
		})
		Convey("return pageInfo ", func() {
			pageInfo := pagination.PageInfo(1)
			So(pageInfo, ShouldResemble, PageInfo{
				TotalItems:  1,
				ItemPerPage: PaginationDefaultItemPerPage,
				PageNumber:  1,
			})
		})
	})
}

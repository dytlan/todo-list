package functions

const (
	PaginationDefaultItemPerPage = 15
)

type PageSpec struct {
	PageNumber  int `json:"pageNumber"`
	ItemPerPage int `json:"itemPerPage"`
}

func (ths PageSpec) OffsetValue() int {
	if ths.PageNumber == 0 {
		return 0
	}
	return (ths.PageNumber - 1) * ths.ItemPerPage
}

func (ths PageSpec) LimitValue() int {
	if ths.ItemPerPage == 0 {
		return PaginationDefaultItemPerPage
	}
	return ths.ItemPerPage
}

func (ths PageSpec) PageInfo(totalRecord int64) PageInfo {
	if ths.ItemPerPage == 0 {
		ths.ItemPerPage = PaginationDefaultItemPerPage
	}
	if ths.PageNumber == 0 {
		ths.PageNumber = 1
	}
	return PageInfo{
		TotalItems:  totalRecord,
		ItemPerPage: ths.ItemPerPage,
		PageNumber:  ths.PageNumber,
	}
}

type PageInfo struct {
	TotalItems  int64 `json:"totalItems"`
	ItemPerPage int   `json:"itemPerPage"`
	PageNumber  int   `json:"pageNumber"`
}

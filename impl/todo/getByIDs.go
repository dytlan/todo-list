package todo

import (
	"github.com/dytlan/moonlay-todo-list/engine/response"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"github.com/dytlan/moonlay-todo-list/impl/functions"
	"github.com/labstack/echo/v4"
)

func (ths *service) GetByIDs(c echo.Context) error {
	var getByIDsReq GetByIDsReq
	if err := functions.BindAndValidate(c, &getByIDsReq); err != nil {
		return err
	}
	getByIDsRes, err := ths.accessor.GetByIDs(c, accessor.GetByIDsSpec{
		IDs:    getByIDsReq.IDs,
		Filter: getByIDsReq.Filter,
	})
	if err != nil {
		return err
	}
	res := newGetByIDsDataOutput(getByIDsRes)
	return response.NewSuccessResponse(c, res)
}

func newGetByIDsDataOutput(data []accessor.ToDoDB) []*GetByIDsData {
	result := make([]*GetByIDsData, 0)
	for _, d := range data {
		res := &GetByIDsData{
			ToDoDB:  d,
			SubData: nil,
		}
		var isFound bool
		for _, root := range result {
			parent := findByID(root, d.MasterID)
			if parent != nil {
				parent.SubData = append(parent.SubData, res)
				isFound = true
				break
			}
		}
		if !isFound {
			result = append(result, res)
		}
	}
	return result
}

func findByID(root *GetByIDsData, ID uint) *GetByIDsData {
	queue := make([]*GetByIDsData, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.ID == ID {
			return nextUp
		}
		if len(nextUp.SubData) > 0 {
			for _, child := range nextUp.SubData {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

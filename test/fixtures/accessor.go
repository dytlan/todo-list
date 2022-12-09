package fixtures

import (
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"gorm.io/gorm"
	"time"
)

var trueStatement, falseStatement = true, false

var DummyAccessorToDo = accessor.ToDoDB{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now().Round(0),
		UpdatedAt: time.Now().Round(0),
	},
	Title:       "title",
	Description: "description",
	File:        "file",
	Level:       1,
	IsDone:      &falseStatement,
	MasterID:    0,
	RootID:      0,
}

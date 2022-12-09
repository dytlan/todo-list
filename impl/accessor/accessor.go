package accessor

import "gorm.io/gorm"

type Accessor interface {
	ToDoAccessor
}

type accessor struct {
	db *gorm.DB
}

func NewAccessor(db *gorm.DB) Accessor {
	return &accessor{
		db: db,
	}
}

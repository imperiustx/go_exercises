package categorymodel

import "github.com/imperiustx/go_excercises/common"

type CategoryUpdate struct {
	Name        string        `json:"name" gorm:"not null"`
	Description string        `json:"description"`
	Icon        *common.Image `json:"icon"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

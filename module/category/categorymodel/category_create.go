package categorymodel

import "github.com/imperiustx/go_excercises/common"

type CategoryCreate struct {
	common.SQLModelCreate `json:",inline"`
	Name                  string        `json:"name" `
	Description           string        `json:"description"`
	Icon                  *common.Image `json:"icon"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

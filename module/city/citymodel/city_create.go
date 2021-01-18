package citymodel

import "github.com/imperiustx/go_excercises/common"

type CityCreate struct {
	common.SQLModelCreate `json:",inline"`
	Title                 string `json:"title"`
}

func (CityCreate) TableName() string {
	return City{}.TableName()
}

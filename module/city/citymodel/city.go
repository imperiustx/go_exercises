package citymodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "City"

// City model
type City struct {
	common.SQLModel
	Title string `json:"title"`
}

// TableName table name
func (City) TableName() string {
	return "cities"
}

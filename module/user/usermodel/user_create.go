package usermodel

import "github.com/imperiustx/go_excercises/common"

type UserCreate struct {
	common.SQLModelCreate `json:",inline"`
	Email                 string `json:"email"`
	Password              string `json:"password"`
	Salt                  string `json:"-"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Phone                 string `json:"phone"`
	Role                  string `json:"-"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

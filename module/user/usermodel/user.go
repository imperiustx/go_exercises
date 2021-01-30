package usermodel

import (
	"errors"

	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "User"

// User model
type User struct {
	common.SQLModel
	Email      string `json:"email" gorm:"not null"`
	FacebookID string `json:"fb_id"`
	GoogleID   string `json:"gg_id"`
	Password   string `json:"password" gorm:"not null"`
	Salt       string
	FirstName  string        `json:"first_name" gorm:"not null"`
	LastName   string        `json:"last_name" gorm:"not null"`
	Phone      string        `json:"phone"`
	Role       string        `json:"role" gorm:"default:user;not null"`
	Avatar     *common.Image `json:"avatar"`
}

// TableName table name
func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DBTypeUser, 1)
}

func (u *User) GetUserId() int {
	return u.ID
}
func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)

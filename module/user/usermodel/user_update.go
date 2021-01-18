package usermodel

type UserUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

package userstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context, id int) (*usermodel.User, error) {
	db := s.db
	var user usermodel.User

	if err := db.Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

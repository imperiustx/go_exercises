package addressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

func (s *sqlStore) FindAddress(ctx context.Context, id int) (*addressmodel.Address, error) {
	db := s.db
	var address addressmodel.Address

	if err := db.Where("id = ?", id).
		First(&address).Error; err != nil {
		return nil, err
	}

	return &address, nil
}

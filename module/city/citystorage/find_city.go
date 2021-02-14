package citystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

func (s *sqlStore) FindCity(ctx context.Context, id int) (*citymodel.City, error) {
	db := s.db.Table(citymodel.City{}.TableName())
	var city citymodel.City

	if err := db.Where("id = ?", id).First(&city).Error; err != nil {
		return nil, err
	}

	return &city, nil
}

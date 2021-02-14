package citystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

func (s *sqlStore) UpdateCity(ctx context.Context, id int, data *citymodel.CityUpdate) error {
	db := s.db.Table(data.TableName())

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

package citystorage

import "github.com/imperiustx/go_excercises/module/city/citymodel"

func (s *sqlStore) DeleteCity(id int) error {
	db := s.db.Table(citymodel.City{}.TableName())

	if err := db.Where("id = ?", id).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}

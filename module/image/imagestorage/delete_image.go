package imagestorage

import "github.com/imperiustx/go_excercises/module/image/imagemodel"

func (s *sqlStore) DeleteImage(id int) error {
	db := s.db

	if err := db.Table("images").
		Delete(&imagemodel.Image{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

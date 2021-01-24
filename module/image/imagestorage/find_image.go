package imagestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/image/imagemodel"
)

func (s *sqlStore) FindImage(ctx context.Context, id int) (*imagemodel.Image, error) {
	db := s.db
	var image imagemodel.Image

	if err := db.Where(&imagemodel.Image{ID: id}).First(&image).Error; err != nil {
		return nil, err
	}

	return &image, nil
}

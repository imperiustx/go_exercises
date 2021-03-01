package uploadstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
)

func (store *sqlStore) ListImages(
	context context.Context,
	ids []int,
	moreKeys ...string,
) ([]common.Image, error) {
	db := store.db.Table(common.Image{}.TableName())
	var result []common.Image

	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (store *sqlStore) ListBunchOfImages(
	context context.Context,
	ids []int,
	moreKeys ...string,
) (common.Images, error) {
	db := store.db.Table(common.Image{}.TableName())
	var result []common.Image

	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

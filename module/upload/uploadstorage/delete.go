package uploadstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
)

func (store *sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	db := store.db.Table(common.Image{}.TableName())

	if err := db.Where("id in (?)", ids).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}

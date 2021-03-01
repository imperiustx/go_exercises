package uploadstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
)

func (store *sqlStore) CreateImage(context context.Context, data *common.Image) error {
	db := store.db.Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

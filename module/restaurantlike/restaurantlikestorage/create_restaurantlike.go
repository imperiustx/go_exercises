package restaurantlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.RestaurantLikeCreate) error {
	db := s.db.Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

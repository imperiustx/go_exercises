package foodlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

func (s *sqlStore) CreateFoodLike(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error {
	db := s.db.Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

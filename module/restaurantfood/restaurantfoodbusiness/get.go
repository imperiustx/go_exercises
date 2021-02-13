package restaurantfoodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

// GetRestaurantFoodStorage get
type GetRestaurantFoodStorage interface {
	FindRestaurantFood(
		ctx context.Context, 
		rid, fid int,
		moreInfo ...string) (*restaurantfoodmodel.RestaurantFood, error)
}

type getRestaurantFoodBiz struct {
	store GetRestaurantFoodStorage
}

// NewGetRestaurantFoodBiz get
func NewGetRestaurantFoodBiz(store GetRestaurantFoodStorage) *getRestaurantFoodBiz {
	return &getRestaurantFoodBiz{store: store}
}

func (biz *getRestaurantFoodBiz) GetRestaurantFood(
	ctx context.Context, 
	rid, fid int,
	moreInfo ...string) (*restaurantfoodmodel.RestaurantFood, error) {

	restaurantfood, err := biz.store.FindRestaurantFood(ctx, rid, fid)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantfoodmodel.EntityName, err)
	}

	return restaurantfood, nil
}

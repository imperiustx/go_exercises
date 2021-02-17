package restaurantfoodbusiness

import (
	"context"
	"strconv"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

// CreateRestaurantFoodStorage create
type CreateRestaurantFoodStorage interface {
	FindRestaurantFood(
		ctx context.Context,
		rid, fid int,
		moreInfo ...string) (*restaurantfoodmodel.RestaurantFood, error)
	CreateRestaurantFood(
		ctx context.Context,
		data *restaurantfoodmodel.RestaurantFoodCreate) error
}

type createRestaurantFood struct {
	store CreateRestaurantFoodStorage
}

// NewCreateRestaurantFoodBiz create
func NewCreateRestaurantFoodBiz(store CreateRestaurantFoodStorage) *createRestaurantFood {
	return &createRestaurantFood{store: store}
}

func (biz *createRestaurantFood) CreateNewRestaurantFood(
	ctx context.Context,
	data *restaurantfoodmodel.RestaurantFoodCreate) error {

	fid, err := strconv.Atoi(data.FoodID)
	if err != nil {
		return common.ErrCannotConvertGivenData(data.FoodID, err)
	}
	rid, err := strconv.Atoi(data.RestaurantID)
	if err != nil {
		return common.ErrCannotConvertGivenData(data.RestaurantID, err)
	}

	restaurantfood, err := biz.store.FindRestaurantFood(ctx, rid, fid)

	if restaurantfood != nil {
		return common.ErrEntityExisted(restaurantfoodmodel.EntityName, err)
	}

	if err := biz.store.CreateRestaurantFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantfoodmodel.EntityName, err)
	}

	return nil
}

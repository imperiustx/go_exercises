package foodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// CreateFoodStorage create
type CreateFoodStorage interface {
	FindFood(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*foodmodel.Food, error)
	CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
}

type ImgStorage interface {
	ListBunchOfImages(
		context context.Context,
		ids []int,
		moreKeys ...string,
	) (common.Images, error)
	DeleteImages(ctx context.Context, ids []int) error
}

type createFood struct {
	store    CreateFoodStorage
	imgStore ImgStorage
}

// NewCreateFoodBiz create
func NewCreateFoodBiz(store CreateFoodStorage, imgStore ImgStorage) *createFood {
	return &createFood{store: store, imgStore: imgStore}
}

func (biz *createFood) CreateNewFood(ctx context.Context, data *foodmodel.FoodCreate) error {

	imgs, err := biz.imgStore.ListBunchOfImages(ctx, data.ImagesID)
	if err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	if len(imgs) == 0 {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	data.Images = &imgs

	food, err := biz.store.FindFood(ctx, map[string]interface{}{"name": data.Name})
	if food != nil {
		return common.ErrEntityExisted(foodmodel.EntityName, err)
	}

	if err := biz.store.CreateFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	go func() {
		_ = biz.imgStore.DeleteImages(ctx, data.ImagesID)
	}()

	return nil
}

// TODO: fix the logic of find food, this one should be find food by name and restaurant

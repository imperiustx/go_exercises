package restaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// CreateRestaurantStorage create
type CreateRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantmodel.Restaurant, error)
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type ImgStorage interface {
	ListImages(
		context context.Context,
		ids []int,
		moreKeys ...string,
	) ([]common.Image, error)
	DeleteImages(ctx context.Context, ids []int) error
}

type createRestaurant struct {
	store    CreateRestaurantStorage
	imgStore ImgStorage
}

// NewCreateRestaurantBiz create
func NewCreateRestaurantBiz(store CreateRestaurantStorage, imgStore ImgStorage) *createRestaurant {
	return &createRestaurant{store: store, imgStore: imgStore}
}

func (biz *createRestaurant) CreateNewRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	imgs, err := biz.imgStore.ListImages(ctx, []int{data.LogoImgID})
	if err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	if len(imgs) == 0 {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	data.Logo = &imgs[0]

	restaurant, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"name": data.Name})
	if restaurant != nil {
		return common.ErrEntityExisted(restaurantmodel.EntityName, err)
	}

	if err := biz.store.CreateRestaurant(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	go func() {
		_ = biz.imgStore.DeleteImages(ctx, []int{data.LogoImgID})
	}()

	return nil
}

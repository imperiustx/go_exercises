package restaurantlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

// GetRestaurantLikeStorage get
type GetRestaurantLikeStorage interface {
	FindRestaurantLike(
		ctx context.Context,
		uid, rid int,
		moreInfo ...string) (*restaurantlikemodel.RestaurantLike, error)
}

type getRestaurantLikeBiz struct {
	store GetRestaurantLikeStorage
}

// NewGetRestaurantLikeBiz get
func NewGetRestaurantLikeBiz(store GetRestaurantLikeStorage) *getRestaurantLikeBiz {
	return &getRestaurantLikeBiz{store: store}
}

func (biz *getRestaurantLikeBiz) GetRestaurantLike(
	ctx context.Context, 
	uid, rid int, 
	moreInfo ...string) (*restaurantlikemodel.RestaurantLike, error) {
		
	restaurantlike, err := biz.store.FindRestaurantLike(ctx, uid, rid)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantlikemodel.EntityName, err)
	}

	return restaurantlike, nil
}

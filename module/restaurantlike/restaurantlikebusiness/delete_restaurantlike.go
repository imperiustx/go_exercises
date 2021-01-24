package restaurantlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

// DeleteRestaurantLikeStorage delete
type DeleteRestaurantLikeStorage interface {
	FindRestaurantLike(ctx context.Context, uid, fid int) (*restaurantlikemodel.RestaurantLike, error)
	DeleteRestaurantLike(uid, rid int) error
}

type deleteRestaurantLike struct {
	store DeleteRestaurantLikeStorage
}

// NewDeleteRestaurantLikeBiz delete
func NewDeleteRestaurantLikeBiz(store DeleteRestaurantLikeStorage) *deleteRestaurantLike {
	return &deleteRestaurantLike{store: store}
}

func (biz *deleteRestaurantLike) DeleteRestaurantLike(ctx context.Context, uid, rid int) error {
	_, err := biz.store.FindRestaurantLike(ctx, uid, rid)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantlikemodel.EntityName, err)
	}
	// TODO: consider add status to model
	// if restaurantlike.Status == 0 {
	// 	return common.ErrCannotGetEntity(restaurantlikemodel.EntityName, errors.New("restaurantlike not found"))
	// }

	if err := biz.store.DeleteRestaurantLike(uid, rid); err != nil {
		return err
	}

	return nil
}

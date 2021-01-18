package citybusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

// DeleteCityStorage delete
type DeleteCityStorage interface {
	FindCity(ctx context.Context, id int) (*citymodel.City, error)
	DeleteCity(id int) error
}

type deleteCity struct {
	store DeleteCityStorage
}

// NewDeleteCityBiz delete
func NewDeleteCityBiz(store DeleteCityStorage) *deleteCity {
	return &deleteCity{store: store}
}

func (biz *deleteCity) DeleteCity(ctx context.Context, id int) error {
	city, err := biz.store.FindCity(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(citymodel.EntityName, err)
	}

	if city.Status == 0 {
		return common.ErrCannotGetEntity(citymodel.EntityName, errors.New("city not found"))
	}

	if err := biz.store.DeleteCity(id); err != nil {
		return err
	}

	return nil
}

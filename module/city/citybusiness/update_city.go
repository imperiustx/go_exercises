package citybusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

// UpdateCityStorage update
type UpdateCityStorage interface {
	FindCity(ctx context.Context, id int) (*citymodel.City, error)
	UpdateCity(ctx context.Context, id int, data *citymodel.CityUpdate) error
}

type updateCity struct {
	store UpdateCityStorage
}

// NewUpdateCityBiz update
func NewUpdateCityBiz(store UpdateCityStorage) *updateCity {
	return &updateCity{store: store}
}

func (biz *updateCity) UpdateCity(ctx context.Context, id int, data *citymodel.CityUpdate) error {
	city, err := biz.store.FindCity(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(citymodel.EntityName, err)
	}

	if city.Status == 0 {
		return common.ErrCannotGetEntity(citymodel.EntityName, errors.New("city not found"))
	}

	if err := biz.store.UpdateCity(ctx, city.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(citymodel.EntityName, err)
	}

	return nil
}

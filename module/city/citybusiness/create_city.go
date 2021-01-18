package citybusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

// CreateCityStorage create
type CreateCityStorage interface {
	CreateCity(ctx context.Context, data *citymodel.CityCreate) error
}

type createCity struct {
	store CreateCityStorage
}

// NewCreateCityBiz create
func NewCreateCityBiz(store CreateCityStorage) *createCity {
	return &createCity{store: store}
}

func (biz *createCity) CreateNewCity(ctx context.Context, data *citymodel.CityCreate) error {
	if err := biz.store.CreateCity(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(citymodel.EntityName, err)
	}

	return nil
}

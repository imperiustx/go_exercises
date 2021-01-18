package citybusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

// GetCityStorage get
type GetCityStorage interface {
	FindCity(ctx context.Context, id int) (*citymodel.City, error)
}

type getCityBiz struct {
	store GetCityStorage
}

// NewGetCityBiz get
func NewGetCityBiz(store GetCityStorage) *getCityBiz {
	return &getCityBiz{store: store}
}

func (biz *getCityBiz) GetCity(ctx context.Context, id int) (*citymodel.City, error) {
	city, err := biz.store.FindCity(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(citymodel.EntityName, err)
	}

	return city, nil
}

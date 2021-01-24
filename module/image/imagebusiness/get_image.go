package imagebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/image/imagemodel"
)

// GetImageStorage get
type GetImageStorage interface {
	FindImage(ctx context.Context, id int) (*imagemodel.Image, error)
}

type getImageBiz struct {
	store GetImageStorage
}

// NewGetImageBiz get
func NewGetImageBiz(store GetImageStorage) *getImageBiz {
	return &getImageBiz{store: store}
}

func (biz *getImageBiz) GetImage(ctx context.Context, id int) (*imagemodel.Image, error) {
	image, err := biz.store.FindImage(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(imagemodel.EntityName, err)
	}

	return image, nil
}

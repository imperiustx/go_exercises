package imagebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/image/imagemodel"
)

// CreateImageStorage create
type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *imagemodel.ImageCreate) error
}

type createImage struct {
	store CreateImageStorage
}

// NewCreateImageBiz create
func NewCreateImageBiz(store CreateImageStorage) *createImage {
	return &createImage{store: store}
}

func (biz *createImage) CreateNewImage(ctx context.Context, data *imagemodel.ImageCreate) error {
	if err := biz.store.CreateImage(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(imagemodel.EntityName, err)
	}

	return nil
}

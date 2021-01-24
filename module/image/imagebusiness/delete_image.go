package imagebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/image/imagemodel"
)

// DeleteImageStorage delete
type DeleteImageStorage interface {
	FindImage(ctx context.Context, id int) (*imagemodel.Image, error)
	DeleteImage(id int) error
}

type deleteImage struct {
	store DeleteImageStorage
}

// NewDeleteImageBiz delete
func NewDeleteImageBiz(store DeleteImageStorage) *deleteImage {
	return &deleteImage{store: store}
}

func (biz *deleteImage) DeleteImage(ctx context.Context, id int) error {
	_, err := biz.store.FindImage(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(imagemodel.EntityName, err)
	}

	if err := biz.store.DeleteImage(id); err != nil {
		return err
	}

	return nil
}

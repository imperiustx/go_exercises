package uploadprovider

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

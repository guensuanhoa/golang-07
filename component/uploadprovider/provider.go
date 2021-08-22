package uploadprovider

import (
	"context"

	"example.com/g07-food-delivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	GetDomain() string
}

package uploadbiz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	"example.com/g07-food-delivery/common"
)

//
//type CreateImageStorage interface {
//	CreateImage(context context.Context, data *common.Image) error
//}

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

type uploadBiz struct {
	provider UploadProvider
}

func NewUploadBiz(provider UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, errors.New("file is not image")
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                              // "a/b/img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt) // 9129324893248.jpg ten file unit

	// Khong muon cho tang business biet cach minh upload file nhu the nao
	// Nen co cai provider o day
	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, errors.New("cannot save image")
	}

	img.Width = w
	img.Height = h
	//img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	// delete img on S3
	//	return nil, uploadmodel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

// Lay width, height tam hinh. Decode de check coi file co phai la image khong
func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}

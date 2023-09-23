package uploadbusiness

import (
	"context"
	"fmt"
	"food-delivery/common"
	"food-delivery/component/uploadprovider"
	"food-delivery/module/upload/uploadmodel"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

//type CreateImageStorage interface {
//	CreateImage(ctx context.Context, data *common.Image) error
//}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	//imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	//fileBytes := bytes.NewBuffer(data)

	//w, h, err := getImageDimension(fileBytes)

	//if err != nil {
	//	return nil, uploadmodel.ErrFileIsNotImage(err)
	//}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)

	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	//img.Width = w
	//img.Height = h
	//img.Extension = fileExt

	return img, nil

}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
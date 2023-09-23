package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"food-delivery/common"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type s3Provider struct {
	bucketName      string
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	region          string
	session         *minio.Client
}

func NewS3Provider(bucketName string, endpoint string, accessKeyID string, secretAccessKey string, region string) *s3Provider {
	provider := &s3Provider{
		bucketName:      bucketName,
		endpoint:        endpoint,
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		region:          region,
	}
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}
	provider.session = minioClient
	return provider
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	provider.MakeBucket(ctx)
	fileByte := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := provider.session.PutObject(ctx, provider.bucketName, dst, fileByte, int64(len(data)), minio.PutObjectOptions{ContentType: fileType})

	if err != nil {
		return nil, err
	}

	//endpointURL := provider.session.EndpointURL()

	img := &common.Image{
		Url: fmt.Sprintf("%s/%s/%s", provider.endpoint, provider.bucketName, dst),
		//Url:       endpointURL.ResolveReference(&url.URL{Path: provider.bucketName + "/" + dst}).String(),
		CloudName: "s3",
	}
	return img, nil
}

func (provider *s3Provider) MakeBucket(ctx context.Context) {
	err := provider.session.MakeBucket(ctx, provider.bucketName, minio.MakeBucketOptions{Region: provider.region})
	if err != nil {
		exists, errBucketExists := provider.session.BucketExists(ctx, provider.bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", provider.bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", provider.bucketName)
	}
}

func buildObjectURL(endpoint, bucket, objectName string) string {
	u := url.URL{
		Scheme: "http",
		Host:   endpoint,
		Path:   strings.Join([]string{bucket, objectName}, "/"),
	}
	return u.String()
}

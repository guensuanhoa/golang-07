package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/g07-food-delivery/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			""),             // Token can be ignore
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	//req, _ := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
	//	Bucket: aws.String(provider.bucketName),
	//	Key:    aws.String(dst),
	//	ACL:    aws.String("private"),
	//})
	//
	//req.Presign(time.Second * 5) // URL

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       dst,
		CloudName: "s3",
	}

	return img, nil
}

func (provider *s3Provider) GetUploadPresignedURL(ctx context.Context) string {
	req, _ := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(provider.bucketName),
		Key:    aws.String(fmt.Sprintf("img/%d", time.Now().UnixNano())),
		ACL:    aws.String("private"),
	})
	//
	url, _ := req.Presign(time.Second * 60)

	return url
}

func (provider *s3Provider) GetDomain() string { return provider.domain }

package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/s-beats/go-cms/env"
)

func NewS3Uploader() *S3Uploader {
	var c client.ConfigProvider
	if env.Envcode() == "local" {
		// https://docs.min.io/docs/how-to-use-aws-sdk-for-go-with-minio-server.html
		c = session.Must(
			session.NewSession(
				&aws.Config{
					Credentials:      credentials.NewStaticCredentials("user", "password", ""),
					Endpoint:         aws.String("http://localhost:9000"),
					Region:           aws.String("ap-northeast-1"),
					DisableSSL:       aws.Bool(true),
					S3ForcePathStyle: aws.Bool(true),
				}))
	} else {
		c = session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")}))
	}
	s3 := s3manager.NewUploader(c)
	return &S3Uploader{s3}
}

type S3Uploader struct {
	*s3manager.Uploader
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	endpoint := "localhost:9000"
	accessKeyID := "user"
	secretAccessKey := "password"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	bucket := "test"

	exists, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		fmt.Printf("failed to verify if bucket exists bucket: %s \n", err)
		return
	}

	if exists {
		if err := minioClient.RemoveBucket(ctx, bucket); err != nil {
			fmt.Printf("failed to remove bucket: %s \n", err)
			return
		}
	}

	if err := minioClient.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{Region: "ap-northeast-1"}); err != nil {
		fmt.Printf("failed to make bucket: %s \n", err)
		return
	}

	fmt.Println("Successfully created bucket.")
}

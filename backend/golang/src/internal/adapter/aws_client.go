package adapter

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type awsClient struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
}

func NewAwsClient(cfg aws.Config, bucketName string) *awsClient {
	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	return &awsClient{client: client, presignClient: presignClient, bucketName: bucketName}
}

func (ac *awsClient) CheckBucket() error {
	output, err := ac.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(ac.bucketName),
	})
	if err != nil {
		log.Println("failed to check bucket of ", ac.bucketName, " :", err)
		return err
	}

	log.Println("first page results")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), *object.Size)
	}

	return nil
}

func (ac *awsClient) UploadImage(ctx context.Context, key string) error {
	imageFile, err := os.Open("test_image.png")
	if err != nil {
		log.Println("error upload image to S3", err)
		return err
	}
	defer imageFile.Close()

	_, err = ac.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(ac.bucketName),
		Key:    aws.String(key),
		Body:   imageFile,
	})
	if err != nil {
		log.Println("error put object: s", err)
		return err
	}

	err = s3.NewObjectExistsWaiter(ac.client).Wait(
		ctx, &s3.HeadObjectInput{Bucket: aws.String(ac.bucketName), Key: aws.String(key)}, time.Minute)
	if err != nil {
		log.Printf("Failed attempt to wait for object %s to exist.\n", key)
	}
	return nil
}

func (ac *awsClient) GetObject(
	ctx context.Context,
	objectKey string,
	lifetimeSecs int64) (*v4.PresignedHTTPRequest, error) {
	request, err := ac.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(ac.bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Println("cannot get presigned")
		return nil, err
	}
	return request, err
}

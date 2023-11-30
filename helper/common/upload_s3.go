package common

import (
	// "html/template"

	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string
var MyBucket string
var filepath string

func ConnectAws() *session.Session {
	err := LoadENV()
	if err != nil {
		log.Fatalln(err)
	}

	AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	MyRegion = os.Getenv("AWS_DEFAULT_REGION")
	AWS_URL := os.Getenv("AWS_URL")

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
			Endpoint:         aws.String(AWS_URL),
			S3ForcePathStyle: aws.Bool(false),
		})

	if err != nil {
		panic(err)
	}

	return sess
}

func UploadFileToS3(sess *session.Session, file *multipart.File, filename string) (string, error) {
	err := LoadENV()
	if err != nil {
		log.Fatalln(err)
	}

	MyBucket = os.Getenv("AWS_BUCKET")
	filenameSlug := ConvertKebabCase(filename)
	filepath = os.Getenv("AWS_OBJECT_URL") + filenameSlug

	// Upload
	object := s3.PutObjectInput{
		Bucket: aws.String(MyBucket),
		Key:    aws.String(filenameSlug),
		Body:   *file,
		ACL:    aws.String("public-read"),
	}

	s3Client := s3.New(sess)
	_, err = s3Client.PutObject(&object)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func DeleteFileFromS3(sess *session.Session, filename string) error {
	err := LoadENV()
	if err != nil {
		log.Fatalln(err)
	}

	MyBucket = os.Getenv("AWS_BUCKET")

	splitedFilename := strings.Split(filename, "/")
	filename = splitedFilename[len(splitedFilename)-1]

	// Delete
	object := s3.DeleteObjectInput{
		Bucket: aws.String(MyBucket),
		Key:    aws.String(filename),
	}

	s3Client := s3.New(sess)
	_, err = s3Client.DeleteObject(&object)
	if err != nil {
		return err
	}

	return nil
}

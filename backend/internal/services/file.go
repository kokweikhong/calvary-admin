package services

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileService interface {
	UploadFile(fileBytes []byte, filename, saveDir string) (string, error)
	// DownloadFile(filename string) ([]byte, error)
	// DeleteFile(filename string) error
}

type fileService struct {
}

func NewFileService() FileService {
	return &fileService{}
}

func (f *fileService) UploadFile(fileBytes []byte, filename, saveDir string) (string, error) {
	key := os.Getenv("DO_SPACES_KEY")
	secret := os.Getenv("DO_SPACES_SECRET")
	endpoint := "https://calvarycarpentry-cloud-storage.sgp1.digitaloceanspaces.com"
	region := "sgp1"
	session := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(endpoint),
		Region:      aws.String(region),
	}))

	uploader := s3manager.NewUploader(session)

	filename = filepath.Join("calvary-admin", saveDir, filename)

	reader := bytes.NewReader(fileBytes)

	bucket := "calvarycarpentry-cloud-storage"

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   reader,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}

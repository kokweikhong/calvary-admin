package services

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	endpoint := "https://sgp1.digitaloceanspaces.com"
	region := "sgp1"
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(endpoint),
		Region:      aws.String(region),
	}

	session, err := session.NewSession(s3Config)
	if err != nil {
		slog.Error("Error creating session", "message", err.Error())
	}
	s3Client := s3.New(session)

	folder := filepath.Join("calvary-admin", saveDir)
	dst := filepath.Join(folder, filename)
	dst = strings.ReplaceAll(dst, "\\", "/")

	result, err := s3Client.PutObject(&s3.PutObjectInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String("calvarycarpentry-cloud-storage"),
		Key:    aws.String(dst),
		Body:   bytes.NewReader(fileBytes),
	})
	slog.Info("filepath", "file", dst)
	if err != nil {
		slog.Error("Error uploading file", "message", err.Error())
		return "", err
	}

	slog.Info("File uploaded successfully", "message", result)
	return dst, nil
}

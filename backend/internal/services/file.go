package services

type FileService interface {
	UploadFile(file []byte, filename string) (string, error)
	DownloadFile(filename string) ([]byte, error)
	DeleteFile(filename string) error
}



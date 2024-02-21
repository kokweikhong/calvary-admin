package handlers

import (
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/kokweikhong/calvary-admin/backend/internal/services"
	"github.com/kokweikhong/calvary-admin/backend/internal/utils"
)

type FileHandler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
}

type fileHandler struct {
	srv         services.FileService
	jsonHandler utils.JSONHandler
}

func NewFileHandler() FileHandler {
	return &fileHandler{
		srv:         services.NewFileService(),
		jsonHandler: utils.NewJSONHandler(),
	}
}

func (f *fileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Accept", "*/*")
	// multipart form
	slog.Info("UploadFile")

	// 30MB max file size
	if err := r.ParseMultipartForm(30 << 20); err != nil {
		slog.Error("Error parsing multipart form", "error", err.Error())
		f.jsonHandler.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// get the file from the multipart form
	file, handler, err := r.FormFile("file")
	if err != nil {
		f.jsonHandler.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// save the file to disk
	defer file.Close()
	fileBytes := make([]byte, handler.Size)
	if _, err := file.Read(fileBytes); err != nil {
		f.jsonHandler.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	fileName := r.FormValue("filename")
	if fileName == "" {
		fileName = handler.Filename
	} else {
		// concatenate the file extension to the filename
		fileName += "." + filepath.Ext(handler.Filename)
	}

	saveDir := r.FormValue("saveDir")

	// upload the file to the S3 bucket
	url, err := f.srv.UploadFile(fileBytes, fileName, saveDir)
	if err != nil {
		f.jsonHandler.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	f.jsonHandler.WriteJSON(w, http.StatusOK, url)
}

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	IndexFilePath = "index.html"
	FormFileField = "myFile"
	MaxUploadSize = 32 << 20
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, IndexFilePath)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedFile, fileHeader, err := r.FormFile(FormFileField)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	fileData, err := io.ReadAll(uploadedFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedResult, err := service.Convert(string(fileData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05.000000000")
	newFileName := fmt.Sprintf("%s%s", timestamp, fileExtension)

	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	if _, err := newFile.WriteString(convertedResult); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(convertedResult)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

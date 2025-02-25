package storage

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"github.com/google/uuid"
)

type LocalStorage struct {
	UploadDir string
}

func NewLocalStorage(uploadDir string) *LocalStorage {
	return &LocalStorage{
		UploadDir: uploadDir,
	}
}

func (l *LocalStorage) SaveFile(file *multipart.FileHeader) (string, error) {
	if err := l.ValidateFileType(file); err != nil {
		return "", err
	}
	
	if _, err := os.Stat(l.UploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(l.UploadDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext
	filePath := filepath.Join(l.UploadDir, newFileName)
	
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	
	return newFileName, nil
}

func (l *LocalStorage) ValidateFileType(file *multipart.FileHeader) error {
	allowedMimeTypes := []string{"image/jpeg", "image/png", "image/gif"}
	fileMime := file.Header.Get("Content-Type")
	
	for _, mimeType := range allowedMimeTypes {
		if fileMime == mimeType {
			return nil
		}
	}
	
	return errors.New("invalid file type, only image files are allowed")
}
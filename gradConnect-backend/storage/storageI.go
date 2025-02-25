package storage

import "mime/multipart"

type Storage interface{
	SaveFile(file *multipart.FileHeader)(string, error)
	ValidateFileType(file *multipart.FileHeader)error
}
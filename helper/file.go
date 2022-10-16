package helper

import (
	"mime/multipart"
	"net/http"
)

const IMAGE_PNG = "image/png"
const IMAGE_JPEG = "image/jpeg"
const IMAGE_JPG = "image/jpg"

func DetectFileType(file multipart.File) string {
	defer file.Close()
	buf := make([]byte, 512)
	_, err := file.Read(buf)

	if err != nil {
		return ""
	}

	return http.DetectContentType(buf)
}

func IsImage(file multipart.File) bool {
	contentType := DetectFileType(file)

	return contentType == IMAGE_PNG || contentType == IMAGE_JPEG || contentType == IMAGE_JPG
}

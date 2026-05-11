package utility

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessImageUpload(c *gin.Context, field string) ([]byte, string, string, error) {
	file, err := c.FormFile(field)
	if err != nil {
		return nil, "", "", fmt.Errorf("gambar tidak boleh kosong")
	}

	src, err := file.Open()
	if err != nil {
		return nil, "", "", fmt.Errorf("gagal membuka gambar yang diunggah")
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return nil, "", "", fmt.Errorf("gagal membaca gambar yang diunggah")
	}

	contentType := file.Header.Get("Content-Type")
	ext := strings.ToLower(filepath.Ext(file.Filename))

	if ext == ".heic" || ext == ".heif" {
		return nil, "", "", fmt.Errorf("format HEIC/HEIF belum didukung, gunakan JPG, PNG, atau WEBP")
	}

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/webp": true,
	}

	if !allowedTypes[contentType] {
		return nil, "", "", fmt.Errorf("format gambar harus JPG, PNG, atau WEBP")
	}

	if ext == "" {
		ext = ".jpg"
	}

	objectPath := fmt.Sprintf("%d/%02d/%s%s",
		time.Now().Year(),
		time.Now().Month(),
		uuid.NewString(),
		ext,
	)

	return fileBytes, objectPath, contentType, nil
}
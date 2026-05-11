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

func ProcessImageUploadKepala(c *gin.Context, field string) ([]byte, string, string, error) {
	file, err := c.FormFile(field)
	if err != nil {
		return nil, "", "", fmt.Errorf("image is required")
	}

	src, err := file.Open()
	if err != nil {
		return nil, "", "", fmt.Errorf("failed to open uploaded file")
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return nil, "", "", fmt.Errorf("failed to read uploaded file")
	}

	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return nil, "", "", fmt.Errorf("file must be an image")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		ext = ".jpg"
	}

	objectPath := fmt.Sprintf("news/%d/%02d/%s%s",
		time.Now().Year(),
		time.Now().Month(),
		uuid.NewString(),
		ext,
	)

	return fileBytes, objectPath, contentType, nil
}
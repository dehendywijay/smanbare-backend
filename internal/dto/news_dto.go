package dto

import (
	"mime/multipart"
)
type NewsRequest struct {
	Title     string `form:"title" binding:"required"`
	Content   string `form:"content" binding:"required"`
	Thumbnail *multipart.FileHeader `form:"image" binding:"required"`
	Kategori  string `form:"kategori"`
}

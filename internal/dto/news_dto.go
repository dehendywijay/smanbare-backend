package dto

import (
	"mime/multipart"
)
type NewsRequest struct {
	Title     string `form:"title" binding:"required"`
	Content   string `form:"content" binding:"required"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
	Kategori  string `form:"kategori"`
}

type EditNewsRequest struct {
	Title     string `form:"title" binding:"required"`
	Content   string `form:"content" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
	Kategori  string `form:"kategori"`
}

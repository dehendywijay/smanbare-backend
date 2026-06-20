package dto

import (
	"mime/multipart"
)
type NewsRequest struct {
	Title     string `form:"title" validate:"required"`
	Content   string `form:"content" validate:"required"`
	Image     *multipart.FileHeader `form:"image" validate:"required"`
	Kategori  string `form:"kategori"`
}

type EditNewsRequest struct {
	Title     string `form:"title" validate:"required"`
	Content   string `form:"content" validate:"required"`
	Image     *multipart.FileHeader `form:"image"`
	Kategori  string `form:"kategori"`
}

package dto

import "mime/multipart"

type AlumniCreateRequest struct {
	Nama        string `form:"nama" validate:"required"`
	Foto        *multipart.FileHeader `form:"foto" validate:"required"`
	Universitas string `form:"universitas" validate:"required"`
	Tahun       string `form:"tahun" validate:"required"`
}

type AlumniEditRequest struct {
	Nama        string `form:"nama" validate:"required"`
	Foto        *multipart.FileHeader `form:"foto"`
	Universitas string `form:"universitas" validate:"required"`
	Tahun       string `form:"tahun" validate:"required"`
}
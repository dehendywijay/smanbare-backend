package dto

import "mime/multipart"

type AlumniCreateRequest struct {
	Nama        string `form:"nama" binding:"required"`
	Foto        *multipart.FileHeader `form:"foto" binding:"required"`
	Universitas string `form:"universitas" binding:"required"`
	Tahun       string `form:"tahun" binding:"required"`
}
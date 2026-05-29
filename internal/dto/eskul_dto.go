package dto

import "mime/multipart"

type CreateEskulRequest struct {
	Nama     string `form:"nama" binding:"required"`
	Pembina  string `form:"pembina" binding:"required"`
	Jadwal   string `form:"jadwal" binding:"required"`
	Prestasi string `form:"prestasi" binding:"required"`
	Tujuan   string `form:"tujuan" binding:"required"`
	Foto     *multipart.FileHeader `form:"foto" binding:"required"`
	Slug     string `form:"slug" binding:"omitempty"`
}

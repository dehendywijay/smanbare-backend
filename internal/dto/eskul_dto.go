package dto

import "mime/multipart"

type CreateEskulRequest struct {
	Nama     string `form:"nama" validate:"required"`
	Pembina  string `form:"pembina" validate:"required"`
	Jadwal   string `form:"jadwal" validate:"required"`
	Prestasi string `form:"prestasi" validate:"required"`
	Tujuan   string `form:"tujuan" validate:"required"`
	Foto     *multipart.FileHeader `form:"foto" validate:"required"`
	Slug     string `form:"slug" validate:"omitempty"`
}

type EditEskulRequest struct {
	Nama     string `form:"nama" validate:"required"`
	Pembina  string `form:"pembina" validate:"required"`
	Jadwal   string `form:"jadwal" validate:"required"`
	Prestasi string `form:"prestasi" validate:"required"`
	Tujuan   string `form:"tujuan" validate:"required"`
	Foto     *multipart.FileHeader `form:"foto"`
	Slug     string `form:"slug" validate:"omitempty"`
}

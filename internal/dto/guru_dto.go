package dto 

import "mime/multipart"

type KepalaSekolahRequest struct {
	Name    string `form:"name" validate:"required"`
	Content string `form:"content" validate:"required"`
	Foto   *multipart.FileHeader `form:"foto" `
}


type GuruRequest struct {
	Nama     string `form:"nama" validate:"required"`
	Jabatan  string `form:"jabatan" validate:"required"`
	Foto     *multipart.FileHeader `form:"foto" validate:"required"`
}

type EditGuruRequest struct {
	Nama     string `form:"nama" validate:"required"`
	Jabatan  string `form:"jabatan" validate:"required"`
	Foto     *multipart.FileHeader `form:"foto"`
}
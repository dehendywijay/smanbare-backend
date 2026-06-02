package dto 

import "mime/multipart"

type KepalaSekolahRequest struct {
	Name    string `form:"name" binding:"required"`
	Content string `form:"content" binding:"required"`
	Foto   *multipart.FileHeader `form:"foto" binding:"required"`
}

type GuruRequest struct {
	Nama     string `form:"nama" binding:"required"`
	Jabatan  string `form:"jabatan" binding:"required"`
	Foto     *multipart.FileHeader `form:"foto" binding:"required"`
}

type EditGuruRequest struct {
	Nama     string `form:"nama" binding:"required"`
	Jabatan  string `form:"jabatan" binding:"required"`
	Foto     *multipart.FileHeader `form:"foto"`
}
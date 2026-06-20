package services

import (
	"gin-app/internal/dto"
	"gin-app/internal/models"
	"gin-app/internal/repository"
	prosesimage "gin-app/pkg/ProsesImage"
	"gin-app/pkg/storage"

	"fmt"

)

type EskulService struct {
	r *repository.EskulRepository
}

func NewEskulService(r *repository.EskulRepository) *EskulService {
	return &EskulService{
		r: r,
	}
}

func (s *EskulService) CreateEskul(eskul dto.CreateEskulRequest) error {
	fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(eskul.Foto)
	if err != nil {
		return fmt.Errorf("Gagal memproses gambar: %v", err)
	}

	publicURL, err := storage.UploadToSupabase("eskul", objectPath, contentType, fileBytes)
	if err != nil {
		return fmt.Errorf("Gagal mengunggah gambar: %v", err)
	}

	eskulUpdated := models.Eskul{
		Nama:     eskul.Nama,
		Pembina:  eskul.Pembina,
		Jadwal:   eskul.Jadwal,
		Prestasi: eskul.Prestasi,
		Foto:     publicURL,
		Tujuan:   eskul.Tujuan,
		Slug:     eskul.Slug,
	}

	err = s.r.CreateEskul(eskulUpdated)
	if err != nil {
		return fmt.Errorf("Gagal membuat data eskul: %v", err)
	}

	return nil
}

func (s *EskulService) GetEskul() (*[]models.Eskul, error) {
	data, err := s.r.GetEskul()
	if err != nil {
		return nil, fmt.Errorf("Gagal mendapatkan data eskul: %v", err)
	}
	return data, nil
}

func (s *EskulService) GetEskulByID(slug string) (models.Eskul, error) {
	data, err := s.r.GetEskulByID(slug)
	if err != nil {
		return data, fmt.Errorf("Gagal mendapatkan data eskul: %v", err)
	}
	return data, nil
}

func (s *EskulService) EditEskul(slug string, updatedEskul dto.EditEskulRequest)  error {
	eskul := models.Eskul{
		Nama:     updatedEskul.Nama,
		Pembina:  updatedEskul.Pembina,
		Jadwal:   updatedEskul.Jadwal,
		Prestasi: updatedEskul.Prestasi,
		Tujuan:   updatedEskul.Tujuan,
		Slug:     updatedEskul.Slug,
	}

	if updatedEskul.Foto != nil {
		oldObjectPath, err := s.r.GetFotoEskul(slug)
		if err != nil {
			return fmt.Errorf("Gagal mendapatkan data eskul: %v", err)
		}

		oldFoto := prosesimage.ExtractObjectPath(oldObjectPath, "eskul")

		fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(updatedEskul.Foto)
		if err != nil {
			return fmt.Errorf("Gagal memproses gambar: %v", err)
		}

		publicURL, err := storage.UpdateSupabaseFile("eskul", oldFoto, objectPath, contentType, fileBytes)
		if err != nil {
			return fmt.Errorf("Gagal memperbarui file: %v", err)
		}

		eskul.Foto = publicURL
	}

	err := s.r.EditEskul(slug, eskul)
	if err != nil {
		return fmt.Errorf("Gagal memperbarui data eskul: %v", err)
	}
	return nil
}

func (s *EskulService) DeleteEskul(slug string) error {
	foto, err := s.r.GetFotoEskul(slug)
	if err != nil {
		return fmt.Errorf("Gagal mendapatkan mengapus data eskul: %v", err)
	}

	if foto != "" {
		fotopath := prosesimage.ExtractObjectPath(foto, "eskul")
		err = storage.DeleteFromSupabase("eskul", fotopath)
		if err != nil {
			return fmt.Errorf("Gagal menghapus foto: %v", err)
		}
	}

	err = s.r.DeleteEskul(slug)
	if err != nil {
		return fmt.Errorf("Gagal menghapus data eskul: %v", err)
	}
	return nil
}



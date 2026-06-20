package services

import (
	"fmt"
	"gin-app/internal/dto"
	"gin-app/internal/models"
	"gin-app/internal/repository"
	prosesimage "gin-app/pkg/ProsesImage"
	"gin-app/pkg/storage"

)

type AlumniServices struct {
	r *repository.AlumniRepository
}

func NewAlumniService(r *repository.AlumniRepository) *AlumniServices {
	return &AlumniServices{
		r: r,
	}
}

func (s *AlumniServices) GetAllAlumni() (*[]models.Alumni, error) {
	return s.r.GetAllAlumni()
}

func (s *AlumniServices) CreateAlumni(alumni *dto.AlumniCreateRequest) error {
	fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload( alumni.Foto)
	if err != nil {
		return fmt.Errorf("Failed to process image: %w", err)
	}

	publicURL, err := storage.UploadToSupabase("alumni", objectPath, contentType, fileBytes)
	if err != nil {
		return fmt.Errorf("Failed to upload image: %w", err)
	}

	alumniModel := models.Alumni{
		Nama:        alumni.Nama,
		Foto:        publicURL,
		Universitas: alumni.Universitas,
		Tahun:       alumni.Tahun,
	}

	err = s.r.CreateAlumni(alumniModel)
	if err != nil {
		return fmt.Errorf("Failed to create alumni: %w", err)
	}

	return err
}

func (s *AlumniServices) UpdateAlumni(id string, alumni dto.AlumniEditRequest) error {
	alumnimodels := models.Alumni{
		Nama:        alumni.Nama,
		Universitas: alumni.Universitas,
		Tahun:       alumni.Tahun,
	}
	if alumni.Foto != nil {
		oldObjectPath, err := s.r.GetFotoAlumni(id)
		if err != nil {
			return fmt.Errorf("Gagal mengambil data foto lama: %w", err)
		} 

		oldFoto := prosesimage.ExtractObjectPath(oldObjectPath, "alumni")

		fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(alumni.Foto)
		if err != nil {
			return fmt.Errorf("Gagal memproses gambar: %w", err)
		}

		publicURL, err := storage.UpdateSupabaseFile("alumni", oldFoto, objectPath, contentType, fileBytes)
		if err != nil {
			return fmt.Errorf("Gagal mengupload file baru: %w", err)
		}
		alumnimodels.Foto = publicURL
	}

	err := s.r.UpdateAlumni(id, &alumnimodels)
	if err != nil {
		return fmt.Errorf("Gagal memperbarui data alumni: %w", err)
	}
	return nil
}

func (s *AlumniServices) DeleteAlumni(id string) error {
	foto, err := s.r.GetFotoAlumni(id)
	if err != nil {
		return fmt.Errorf("Gagal mengambil data foto alumni: %w", err)
	}

	if foto != "" {
		fotopath := prosesimage.ExtractObjectPath(foto, "alumni")
		err = storage.DeleteFromSupabase("alumni", fotopath)
		if err != nil {
			return fmt.Errorf("Gagal menghapus foto dari Supabase: %w", err)
		}
	}

	err = s.r.DeleteAlumni(id)
	if err != nil {
		return fmt.Errorf("Gagal menghapus data alumni: %w", err)
	}

	return nil
}

package services

import (
	"fmt"
	"gin-app/internal/dto"
	"gin-app/internal/models"
	"gin-app/internal/repository"

	prosesimage "gin-app/pkg/ProsesImage"
	"gin-app/pkg/storage"

	"github.com/gin-gonic/gin"
)

type GuruService struct {
	r *repository.GuruRepository
}

func NewGuruService(r *repository.GuruRepository) *GuruService {
	return &GuruService{
		r: r,
	}
}

func (s *GuruService) CreateGuru(guru dto.GuruRequest, c *gin.Context) error {
	fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(c, guru.Image)
	if err != nil {
		return fmt.Errorf("gagal memproses gambar : %s", err.Error())
	}

	publicURL, err := storage.UploadToSupabase("guru", objectPath, contentType, fileBytes)
	if err != nil {
		return fmt.Errorf("gagal mengunggah gambar : %s", err.Error())
	}

	newGuru := models.Guru{
		Nama:    guru.Nama,
		Jabatan: guru.Jabatan,
		Foto:    publicURL,
	}

	err = s.r.CreateGuru(newGuru)
	if err != nil {
		return fmt.Errorf("gagal menyimpan data guru: %s", err.Error())
	}

	return nil
}

func (s *GuruService) GetGuru() ([]models.Guru, error) {
	return s.r.GetGuru()
}

func (s *GuruService) EditGuru(c *gin.Context, id string, updatedGuru dto.GuruRequest) error {
	guru := models.Guru{
		Nama:    updatedGuru.Nama,
		Jabatan: updatedGuru.Jabatan,
	}

	if updatedGuru.Image != nil {
		oldObjectPath, err := s.r.GetFotoGuru(id)
		if err != nil {
			return fmt.Errorf("gagal mendapatkan foto lama: %s", err.Error())
		}

		oldFoto := prosesimage.ExtractObjectPath(oldObjectPath, "guru")

		fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(c, updatedGuru.Image)
		if err != nil {
			return fmt.Errorf("gagal memproses gambar: %s", err.Error())
		}

		publicURL, err := storage.UpdateSupabaseFile("guru", oldFoto, objectPath, contentType, fileBytes)
		if err != nil {
			return fmt.Errorf("gagal mengupdate gambar: %s", err.Error())
		}

		guru.Foto = publicURL
	}

	err := s.r.EditGuru(id, guru)
	if err != nil {
		return fmt.Errorf("gagal mengupdate data guru: %s", err.Error())
	}
	return nil
}

func (s *GuruService) DeleteGuru(id string) error {
	foto, err := s.r.GetFotoGuru(id)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan data guru: %s", err.Error())
	}

	if foto != "" {
		fotopath := prosesimage.ExtractObjectPath(foto, "guru")
		err = storage.DeleteFromSupabase("guru", fotopath)
		if err != nil {
			return fmt.Errorf("gagal menghapus foto: %s", err.Error())
		}
	}

	err = s.r.DeleteGuru(id)
	if err != nil {
		return fmt.Errorf("gagal menghapus data guru: %s", err.Error())
	}
	return nil
}

func (s *GuruService) EditKepala(id string, updatedKepalaSekolah dto.KepalaSekolahRequest, c *gin.Context) error {
	kepalaSekolah := models.KepalaSekolah{
		Name:    updatedKepalaSekolah.Name,
		Content: updatedKepalaSekolah.Content,
	}
	if updatedKepalaSekolah.Image != nil {
		oldObjectPath, err := s.r.GetFotoKepala(id)
		if err != nil {
			return fmt.Errorf("gagal mendapatkan foto lama: %s", err.Error())
		}

		oldFoto := prosesimage.ExtractObjectPath(oldObjectPath, "kepala")

		fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(c, updatedKepalaSekolah.Image)
		if err != nil {
			return fmt.Errorf("gagal memproses gambar: %s", err.Error())
		}

		publicURL, err := storage.UpdateSupabaseFile("kepala", oldFoto, objectPath, contentType, fileBytes)
		if err != nil {
			return fmt.Errorf("gagal mengupdate gambar: %s", err.Error())
		}

		kepalaSekolah.Foto = publicURL
	}

	err := s.r.EditKepala(id, kepalaSekolah)
	if err != nil {
		return fmt.Errorf("gagal mengupdate data kepala sekolah: %s", err.Error())
	}

	return nil
}


// func CreateKepala(kepalaSekolah models.KepalaSekolah) (models.KepalaSekolah, error) {
// 	result := config.DB.Create(&kepalaSekolah)
// 	return kepalaSekolah, result.Error
// }

func (s *GuruService) GetKepalaByID(id string) (models.KepalaSekolah, error) {
	return s.r.GetKepalaByID(id)
}

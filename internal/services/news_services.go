package services

import (
	"fmt"
	"gin-app/internal/dto"
	"gin-app/internal/models"
	"gin-app/internal/repository"
	prosesimage "gin-app/pkg/ProsesImage"
	"gin-app/pkg/slug"
	"gin-app/pkg/storage"

)

type NewsService struct {
	r *repository.NewsRepository
}

func NewNewsService(r *repository.NewsRepository) *NewsService {
	return &NewsService{
		r: r,
	}
}

func (s *NewsService) CreateNews(data dto.NewsRequest) error {
	fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(data.Image)
	if err != nil {
		return fmt.Errorf("failed to process image upload: %w", err)
	}

	publicURL, err := storage.UploadToSupabase("news_thumbnail", objectPath, contentType, fileBytes)
	if err != nil {
		return fmt.Errorf("failed to upload image to Supabase: %w", err)
	}

	slug := slug.MakeSlug(data.Title)

	news := models.News{
		Title:     data.Title,
		Kategori:  data.Kategori,
		Content:   data.Content,
		Thumbnail: publicURL,
		Slug:      slug,
	}

	err = s.r.CreateNews(news)
	if err != nil {
		return fmt.Errorf("failed to create news: %w", err)
	}
	return nil
}

func (s *NewsService) GetNews() ([]models.News, error) {
	result, err := s.r.GetNews()
	if err != nil {
		return nil, fmt.Errorf("failed to get news: %w", err)
	}
	return result, nil
}

func (s *NewsService) GetNewsByID(slug string) (models.News, error) {
	result, err := s.r.GetNewsByID(slug)
	if err != nil {
		return result, fmt.Errorf("failed to get news by ID: %w", err)
	}

	return result, err
}

func (s *NewsService) UpdateNews(slug string, news dto.EditNewsRequest) error {
	newsUpdated := models.News{
		Title:    news.Title,
		Content:  news.Content,
		Slug:     slug,
		Kategori: news.Kategori,
	}

	
	if news.Image != nil {
		oldObjectPath, err := s.r.GetFotoNews(slug)
		if err != nil {
			return fmt.Errorf("failed to get old image path: %w", err)
		}

		fileBytes, objectPath, contentType, err := prosesimage.ProcessImageUpload(news.Image)
		if err != nil {
			return fmt.Errorf("failed to process image upload: %w", err)
		}

		oldPath := prosesimage.ExtractObjectPath(oldObjectPath, "news_thumbnail")
		publicURL, err := storage.UpdateSupabaseFile("news_thumbnail", oldPath, objectPath, contentType, fileBytes)
		if err != nil {
			return fmt.Errorf("failed to update image in Supabase: %w", err)
		}

		newsUpdated.Thumbnail = publicURL
	}

	err := s.r.UpdateNews(slug, newsUpdated)
	if err != nil {
		return fmt.Errorf("failed to update news: %w", err)
	}
	return nil
}

func (s *NewsService) DeleteNews(slug string) error {
	foto, err := s.r.GetFotoNews(slug)
	if err != nil {
		return fmt.Errorf("failed to get news thumbnail: %w", err)
	}

	fotopath := prosesimage.ExtractObjectPath(foto, "news_thumbnail")

	err = storage.DeleteFromSupabase("news_thumbnail", fotopath)
	if err != nil {
		return fmt.Errorf("failed to delete image from Supabase: %w", err)
	}

	err = s.r.DeleteNews(slug)
	if err != nil {
		return fmt.Errorf("failed to delete news: %w", err)
	}

	return nil
}

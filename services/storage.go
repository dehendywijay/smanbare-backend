package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadToSupabase(bucket, objectPath, contentType string, fileBytes []byte) (string, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	serviceRoleKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

	if supabaseURL == "" || serviceRoleKey == "" {
		return "", fmt.Errorf("supabase env is missing")
	}

	endpoint := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseURL, bucket, objectPath)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+serviceRoleKey)
	req.Header.Set("apikey", serviceRoleKey)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("x-upsert", "false")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("Gagal upload ke Supabase (Status %d): %s", resp.StatusCode, string(body))
	}

	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucket, objectPath)
	return publicURL, nil
}

func DeleteFromSupabase(bucket, objectPath string) error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	serviceRoleKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

	if supabaseURL == "" || serviceRoleKey == "" {
		return fmt.Errorf("supabase env is missing")
	}

	endpoint := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseURL, bucket, objectPath)

	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+serviceRoleKey)
	req.Header.Set("apikey", serviceRoleKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed delete file: %s", string(body))
	}

	return nil
}

func UpdateSupabaseFile(bucket, oldObjectPath, newObjectPath, contentType string, fileBytes []byte) (string, error) {
	if oldObjectPath != "" {
		if err := DeleteFromSupabase(bucket, oldObjectPath); err != nil {
			return "", fmt.Errorf("Foto lama gagal dihapus")
		}
	}

	publicURL, err := UploadToSupabase(bucket, newObjectPath, contentType, fileBytes)
	if err != nil {
		return "", err
	}

	return publicURL, nil
}
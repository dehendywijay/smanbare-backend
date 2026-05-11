package dto

type AdminDTO struct {
	Username string `json:"username" binding:"required"`
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
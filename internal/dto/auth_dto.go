package dto

type AdminDTO struct {
	Username string `json:"username" validate:"required"`
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}
package dto

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type TokenPairResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

package gosquare

import "time"

// TokenRequest is the response for requesting a token
type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	AccessToken  string `json:"access_token"`
}

// TokenResponse is the response for requesting a token
type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresAt   time.Time `json:"expires_at"`
	MerchantID  string    `json:"merchant_id"`
}

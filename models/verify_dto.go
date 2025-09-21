package models

// API'ye gelen istek
type VerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// API'den dönen yanıt
type VerifyResponse struct {
	Message string `json:"message"`
}

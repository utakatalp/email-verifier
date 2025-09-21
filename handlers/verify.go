package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utakatalp/email-verifier/models"
	"github.com/utakatalp/email-verifier/services"
)

func VerifyEmail(c *gin.Context) {
	var req models.VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad e-mail pattern"})
		return
	}

	_, err := services.StartVerification(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.VerifyResponse{Message: "Activation email sent"})
}

func ActivateEmail(c *gin.Context) {
	token := c.Query("token")

	email, err := services.CompleteVerification(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully", "email": email})
}

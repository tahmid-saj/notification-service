package routes

import (
	"net/http"
	"notification-service/models"

	"github.com/gin-gonic/gin"
)

// email
func verifyEmail(context *gin.Context) {
	var emailVerificationInput models.EmailVerificationInput

	err := context.ShouldBindJSON(emailVerificationInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	resVerifyEmail, err := models.VerifyEmail(&emailVerificationInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not verify email"})
		return
	}

	context.JSON(http.StatusOK, resVerifyEmail)
}

func sendEmail(context *gin.Context) {
	var emailInput models.EmailInput

	err := context.ShouldBindJSON(emailInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	resSendEmail, err := models.SendEmail(&emailInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not send email"})
		return
	}

	context.JSON(http.StatusOK, resSendEmail)
}
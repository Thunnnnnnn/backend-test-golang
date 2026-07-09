package handlers

import (
	"backend-test-golang/helpers"
	"backend-test-golang/models"
	"backend-test-golang/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := services.Login(&user)
	if err != nil {
		helpers.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{"token": token})
}

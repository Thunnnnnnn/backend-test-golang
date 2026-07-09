package services

import (
	"backend-test-golang/models"
	"backend-test-golang/repositories"
)

func Login(user *models.User) (string, error) {
	return repositories.Login(user)
}

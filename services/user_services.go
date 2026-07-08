package services

import (
	"backend-test-golang/models"
	"backend-test-golang/repositories"
)

func FindAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

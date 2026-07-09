package services

import (
	"backend-test-golang/models"
	"backend-test-golang/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

func FindUserByID(id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &models.User{}, err
	}
	return repositories.FindUserByID(objID)
}

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func UpdateUser(id string, user *models.User) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return repositories.UpdateUser(objID, user)
}

func DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return repositories.DeleteUser(objID)
}

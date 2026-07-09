package repositories

import (
	"backend-test-golang/database"
	"backend-test-golang/helpers"
	"backend-test-golang/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(user *models.User) (string, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		return "", fmt.Errorf("user collection is not initialized")
	}

	var foundUser models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	if !helpers.CheckPasswordHash(user.Password, foundUser.Password) {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := helpers.GenerateJWT(foundUser.ID.Hex(), foundUser)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

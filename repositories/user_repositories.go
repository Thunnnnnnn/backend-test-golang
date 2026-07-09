package repositories

import (
	"backend-test-golang/database"
	"backend-test-golang/helpers"
	"backend-test-golang/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllUsers() ([]models.User, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		return nil, nil
	}

	cursor, err := database.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func FindUserByID(id primitive.ObjectID) (*models.User, error) {
	ctx := context.Background()
	fmt.Println(id)
	if database.UserCollection == nil {
		return nil, nil
	}

	var user models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		return nil, nil
	}

	var user models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	ctx := context.Background()

	find, err := FindUserByEmail(user.Email)
	if find != nil {
		return fmt.Errorf("email already exists")
	}

	if database.UserCollection == nil {
		return nil
	}

	if !helpers.IsValidEmail(user.Email) {
		return fmt.Errorf("invalid email format")
	}

	user.Password, _ = helpers.HashPassword(user.Password)
	user.CreatedAt = time.Now()

	result, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func UpdateUser(id primitive.ObjectID, user *models.User) error {
	ctx := context.Background()

	if database.UserCollection == nil {
		return nil
	}
	existingUser, _ := FindUserByID(id)

	if existingUser == nil {
		return fmt.Errorf("user not found")
	}

	if !helpers.IsValidEmail(user.Email) {
		return fmt.Errorf("invalid email format")
	}

	if user.Password != "" {
		user.Password, _ = helpers.HashPassword(user.Password)
	} else {
		user.Password = existingUser.Password
	}

	user.CreatedAt = existingUser.CreatedAt

	_, err := database.UserCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func DeleteUser(id primitive.ObjectID) error {
	ctx := context.Background()

	if database.UserCollection == nil {
		return nil
	}

	_, err := database.UserCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

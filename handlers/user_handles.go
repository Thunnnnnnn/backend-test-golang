package handlers

import (
	"backend-test-golang/helpers"
	"backend-test-golang/models"
	"backend-test-golang/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers(c *gin.Context) {
	users, err := services.FindAllUsers()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.FindUserByID(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		helpers.Error(c, http.StatusNotFound, "User not found")
		return
	}

	userResponse := UserResponse{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}

	helpers.Success(c, http.StatusOK, userResponse)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.CreateUser(&user)

	if err != nil {
		if err.Error() == "email already exists" {
			helpers.Error(c, http.StatusConflict, err.Error())
			return
		} else if err.Error() == "invalid email format" {
			helpers.Error(c, http.StatusBadRequest, err.Error())
			return
		} else {
			helpers.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	userResponse := UserResponse{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}

	helpers.Success(c, http.StatusCreated, userResponse)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.UpdateUser(id, &user)
	if err != nil {
		if err.Error() == "user not found" {
			helpers.Error(c, http.StatusNotFound, err.Error())
			return
		} else if err.Error() == "invalid email format" {
			helpers.Error(c, http.StatusBadRequest, err.Error())
			return
		} else {
			helpers.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	userResponse := UserResponse{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}

	helpers.Success(c, http.StatusOK, userResponse)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			helpers.Error(c, http.StatusNotFound, err.Error())
			return
		} else {
			helpers.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	helpers.Success(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}

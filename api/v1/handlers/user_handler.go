package handlers

import (
	"net/http"
	"strconv"

	"github.com/bhovdair/go-rest/models"
	"github.com/bhovdair/go-rest/services"
	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetAllUsers handles GET request to retrieve all users
func (h *UserHandler) GetAllUsers(context *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, users)
}

// GetUserByID handles GET request to retrieve a user by ID
func (h *UserHandler) GetUserByID(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

// CreateUser handles POST request to create a new user
func (h *UserHandler) CreateUser(context *gin.Context) {
	var newUser models.User
	if err := context.BindJSON(&newUser); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	newUserID, err := h.userService.CreateUser(newUser)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	newUser.ID = newUserID
	context.IndentedJSON(http.StatusCreated, newUser)
}

// UpdateUser handles PUT request to update a user by ID
func (h *UserHandler) UpdateUser(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedUser models.User
	if err := context.BindJSON(&updatedUser); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	err = h.userService.UpdateUser(uint(id), updatedUser)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	context.IndentedJSON(http.StatusOK, updatedUser)
}

// DeleteUser handles DELETE request to delete a user by ID
func (h *UserHandler) DeleteUser(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

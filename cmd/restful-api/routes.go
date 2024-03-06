package main

import (
	"database/sql"

	"github.com/bhovdair/go-rest/api/v1/handlers"
	"github.com/bhovdair/go-rest/repositories"
	"github.com/bhovdair/go-rest/services"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	// Group routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetAllUsers)
		v1.GET("/users/:id", userHandler.GetUserByID)
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}
}

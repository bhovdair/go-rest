package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/bhovdair/go-rest/api/v1/handlers"
	"github.com/bhovdair/go-rest/repositories"
	"github.com/bhovdair/go-rest/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	router := gin.Default()
	userHandler := handlers.NewUserHandler(userService)

	router.GET("/api/v1/users", userHandler.GetAllUsers)
	router.GET("/api/v1/users/:id", userHandler.GetUserByID)
	router.POST("/api/v1/users", userHandler.CreateUser)
	router.PUT("/api/v1/users/:id", userHandler.UpdateUser)
	router.DELETE("/api/v1/users/:id", userHandler.DeleteUser)

	router.Run("localhost:8080")

	//migrate -database "mysql://root@tcp(localhost:3307)/go-rest-api" -path migrations up
	//migrate -database "mysql://root@tcp(localhost:3307)/go-rest-api" -path migrations down

}

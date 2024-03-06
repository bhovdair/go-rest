package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bhovdair/go-rest/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (uint, error)
	UpdateUser(id uint, user models.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user models.User) (uint, error) {
	query := `
		INSERT INTO users(username, email, password, name)
		VALUES(?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Name)
	if err != nil {
		message := "error when executing query"
		detailedErrorMessage := fmt.Sprintf("%s: %v", message, err.Error())
		log.Printf("Detailed error message: %v", detailedErrorMessage)
		return 0, fmt.Errorf("failed to execute query '%s': %w", query, err)
	}

	newUserID, err := result.LastInsertId()
	if err != nil {
		message := "error retrieving result"
		detailedErrorMessage := fmt.Sprintf("%s: %v", message, err.Error())
		log.Printf("Detailed error message: %v", detailedErrorMessage)
		return 0, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	return uint(newUserID), nil
}

func (r *userRepository) UpdateUser(id uint, user models.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, id)
	return err
}

func (r *userRepository) DeleteUser(id uint) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

package services

import "github.com/bhovdair/go-rest/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (uint, error)
	UpdateUser(id uint, user models.User) error
	DeleteUser(id uint) error
}

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (uint, error)
	UpdateUser(id uint, user models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) CreateUser(user models.User) (uint, error) {
	return s.userRepo.CreateUser(user)
}

func (s *userService) UpdateUser(id uint, user models.User) error {
	return s.userRepo.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}

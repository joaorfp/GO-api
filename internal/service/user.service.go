package service

import (
	"github.com/joaorfp/GO-api/internal/model"
	"github.com/joaorfp/GO-api/internal/repository"
)

type UserService interface {
    GetAllUsers() []model.User
    GetUserById(id int) (model.User, bool)
    CreateUser(user model.User) model.User
    UpdateUserById(id int, user model.User) (model.User, bool)
    DeleteUser(id int) bool
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) GetAllUsers() []model.User {
    return s.repo.GetAll()
}

func (s *userService) GetUserById(id int) (model.User, bool) {
    return s.repo.GetById(id)
}

func (s *userService) CreateUser(user model.User) model.User {
    return s.repo.Create(user)
}

func (s *userService) UpdateUserById(id int, user model.User) (model.User, bool) {
    return s.repo.UpdateById(id, user)
}

func (s *userService) DeleteUser(id int) bool {
    return s.repo.Delete(id)
}

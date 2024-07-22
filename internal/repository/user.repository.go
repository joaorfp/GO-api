package repository

import (
	"github.com/joaorfp/GO-api/internal/model"
)

type UserRepository interface {
    GetAll() []model.User
    GetById(id int) (model.User, bool)
    Create(user model.User) model.User
    UpdateById(id int, user model.User) (model.User, bool)
    Delete(id int) bool
}

type InMemoryUserRepository struct {
    users []model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users: []model.User{},
    }
}

func (r *InMemoryUserRepository) GetAll() []model.User {
    return r.users
}

func (r *InMemoryUserRepository) GetById(id int) (model.User, bool) {
    for _, user := range r.users {
        if user.ID == id {
            return user, true
        }
    }
    return model.User{}, false
}

func (r *InMemoryUserRepository) Create(user model.User) model.User {
    r.users = append(r.users, user)
    return user
}

func (r *InMemoryUserRepository) UpdateById(id int, user model.User) (model.User, bool) {
    for i, u := range r.users {
        if u.ID == id {
            r.users[i] = user
            return user, true
        }
    }
    return model.User{}, false
}

func (r *InMemoryUserRepository) Delete(id int) bool {
    for i, user := range r.users {
        if user.ID == id {
            r.users = append(r.users[:i], r.users[i+1:]...)
            return true
        }
    }
    return false
}
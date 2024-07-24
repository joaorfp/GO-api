package repository

import (
	"github.com/joaorfp/GO-api/internal/model"
)

type UserRepository interface {
    GetAll() []*model.User
    GetById(id int) (*model.User, bool)
    Create(user *model.User) *model.User
    UpdateById(id int, user *model.User) (*model.User, bool)
    Delete(id int) bool
}

// implementação da interface UserRepository. Estocando dados dos usuários localmente usando memória
type InMemoryUserRepository struct {
    users []*model.User
}

// função irá retornar um InMemoryUserRepository que é um tipo. O tipo X será retornado em forma de ponteiro pra variável na memória
// constructor para a estrutura InMemoryUserRepository. Inicializa a estrutura com um slice vazio
func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users: []*model.User{},
    }
}

//  *InMemoryUserRepository é uma instancia da interface de dados. Podendo ser compartilhada no codigo todo. r são os dados em memória
func (r *InMemoryUserRepository) GetAll() []*model.User {
    return r.users
}

func (r *InMemoryUserRepository) GetById(id int) (*model.User, bool) {
    for _, user := range r.users {
        if user.ID == id {
            return user, true
            // retornar user diretamente pois já é um tipo *model.User. Retornar usando &user criaria um ponteiro duplo
        }
    }
    return nil, false
    // nil é exatamente o mesmo que retornar *model.User{} que é um ponteiro para um objeto vazio
}

func (r *InMemoryUserRepository) Create(user *model.User) *model.User {
    r.users = append(r.users, user)
    return user
}

func (r *InMemoryUserRepository) UpdateById(id int, user *model.User) (*model.User, bool) {
    for i, u := range r.users {
        if u.ID == id {
            r.users[i] = user
            return user, true
        }
    }
    return nil, false
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

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaorfp/GO-api/internal/handler"
	"github.com/joaorfp/GO-api/internal/repository"
	"github.com/joaorfp/GO-api/internal/service"
)

func main() {
    // Criação do repositório
    repo := repository.NewInMemoryUserRepository()

    // Criação do serviço
    srv := service.NewUserService(repo)

    // Criação do handler
    hndlr := handler.NewUserHandler(srv)

    // Configuração do roteador
    router := mux.NewRouter()

    router.HandleFunc("/users", hndlr.GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", hndlr.GetUserById).Methods("GET")
    router.HandleFunc("/users", hndlr.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", hndlr.UpdateUserById).Methods("PUT")
    router.HandleFunc("/users/{id}", hndlr.DeleteUser).Methods("DELETE")

    // Inicia o servidor
    fmt.Println("server is on")
    http.ListenAndServe(":8000", router)
}

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joaorfp/GO-api/internal/model"
	"github.com/joaorfp/GO-api/internal/service"
)

type BankHandler struct {
    service service.BankService
}

func NewBankHandler(service service.BankService) *BankHandler {
    return &BankHandler{service: service}
}

func (h *BankHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
    accounts := h.service.GetAllAccounts()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(accounts)
}

func (h *BankHandler) GetAccountByUserId(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userId, err := strconv.Atoi(params["userId"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
        return
    }

    account, found := h.service.GetAccountByUserId(userId)
    if found {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(account)
        return
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "Bank account not found"})
}

func (h *BankHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
    var account model.BankAccount
    _ = json.NewDecoder(r.Body).Decode(&account)
    createdBankAccount := h.service.CreateAccount(&account)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdBankAccount)
}

func (h *BankHandler) UpdateByUserId(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userId, err := strconv.Atoi(params["userId"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
        return
    }

    var account model.BankAccount
    _ = json.NewDecoder(r.Body).Decode(&account)
    updatedAccount, found := h.service.UpdateAccountByUserId(userId, &account)
    if found {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(updatedAccount)
        return
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "Bank account not found"})
}

func (h *BankHandler) DeleteAccountByUserId(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userId, err := strconv.Atoi(params["userId"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
        return
    }

    deleted := h.service.DeleteAccountByUserId(userId)
    if deleted {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"message": "Bank account deleted"})
        return
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "Bank account not found"})
}

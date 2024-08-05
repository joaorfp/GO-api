package model

type BankAccount struct {
    ID             int     `json:"id"`
    Balance        float64 `json:"balance"`
    InvestedBalance float64 `json:"invested_balance"`
    UserID         int     `json:"user_id"` // Foreign key to User
}
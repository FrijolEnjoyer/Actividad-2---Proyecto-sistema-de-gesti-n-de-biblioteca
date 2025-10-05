package models

type LoanRequest struct {
	UserID string `json:"userId"`
	BookID string `json:"bookId"`
}

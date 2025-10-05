package services

import (
	"testing"
	"library/internal/models"
)

func TestBorrowReturnFlow(t *testing.T) {
	s := NewLibraryService()
	s.AddUser(models.User{ID:"u1", Name:"Ana"})
	s.AddBook(models.Book{ID:"b1", Title:"Go", Author:"Gopher"})
	if err := s.Borrow(models.LoanRequest{UserID:"u1", BookID:"b1"}); err != nil { t.Fatalf("borrow: %v", err) }
	// second borrow should fail (not available)
	if err := s.Borrow(models.LoanRequest{UserID:"u1", BookID:"b1"}); err == nil { t.Fatalf("expected error") }
	if err := s.Return(models.LoanRequest{UserID:"u1", BookID:"b1"}); err != nil { t.Fatalf("return: %v", err) }
}

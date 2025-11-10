package services

import (
	"strings"
	"testing"

	"library/internal/models"
)

func TestBorrowReturnFlow(t *testing.T) {
	s := NewLibraryService()
	s.AddUser(models.User{ID: "u1", Name: "Ana"})
	s.AddBook(models.Book{ID: "b1", Title: "Go", Author: "Gopher"})
	if err := s.Borrow(models.LoanRequest{UserID: "u1", BookID: "b1"}); err != nil {
		t.Fatalf("borrow: %v", err)
	}
	// second borrow should fail (not available)
	if err := s.Borrow(models.LoanRequest{UserID: "u1", BookID: "b1"}); err == nil {
		t.Fatalf("expected error")
	}
	if err := s.Return(models.LoanRequest{UserID: "u1", BookID: "b1"}); err != nil {
		t.Fatalf("return: %v", err)
	}
}

func TestBorrowRequiresUserAndBook(t *testing.T) {
	s := NewLibraryService()
	s.AddBook(models.Book{ID: "b1", Title: "Go", Author: "Gopher"})

	if err := s.Borrow(models.LoanRequest{UserID: "missing", BookID: "b1"}); err == nil {
		t.Fatalf("expected error when user does not exist")
	}

	s.AddUser(models.User{ID: "u1", Name: "Ana"})
	if err := s.Borrow(models.LoanRequest{UserID: "u1", BookID: "missing"}); err == nil {
		t.Fatalf("expected error when book does not exist")
	}
}

func TestRemoveBookConstraints(t *testing.T) {
	s := NewLibraryService()
	s.AddUser(models.User{ID: "u1", Name: "Ana"})
	s.AddBook(models.Book{ID: "b1", Title: "Go", Author: "Gopher"})

	if err := s.RemoveBook("missing"); err == nil {
		t.Fatalf("expected error removing missing book")
	}

	if err := s.Borrow(models.LoanRequest{UserID: "u1", BookID: "b1"}); err != nil {
		t.Fatalf("borrow: %v", err)
	}
	if err := s.RemoveBook("b1"); err == nil || !strings.Contains(err.Error(), "loaned") {
		t.Fatalf("expected loaned error removing book, got: %v", err)
	}
	if err := s.Return(models.LoanRequest{UserID: "u1", BookID: "b1"}); err != nil {
		t.Fatalf("return: %v", err)
	}
	if err := s.RemoveBook("b1"); err != nil {
		t.Fatalf("remove after return: %v", err)
	}
	if len(s.ListBooks()) != 0 {
		t.Fatalf("expected no books after removal")
	}
}

func TestRemoveUserConstraints(t *testing.T) {
	s := NewLibraryService()
	s.AddUser(models.User{ID: "u1", Name: "Ana"})
	s.AddBook(models.Book{ID: "b1", Title: "Go", Author: "Gopher"})

	if err := s.RemoveUser("missing"); err == nil {
		t.Fatalf("expected error removing missing user")
	}

	if err := s.Borrow(models.LoanRequest{UserID: "u1", BookID: "b1"}); err != nil {
		t.Fatalf("borrow: %v", err)
	}
	if err := s.RemoveUser("u1"); err == nil || !strings.Contains(err.Error(), "active loans") {
		t.Fatalf("expected active loans error removing user, got: %v", err)
	}
	if err := s.Return(models.LoanRequest{UserID: "u1", BookID: "b1"}); err != nil {
		t.Fatalf("return: %v", err)
	}
	if err := s.RemoveUser("u1"); err != nil {
		t.Fatalf("remove user: %v", err)
	}
	if len(s.ListUsers()) != 0 {
		t.Fatalf("expected no users after removal")
	}
}

func TestSearchBooks(t *testing.T) {
	s := NewLibraryService()
	s.AddBook(models.Book{ID: "b1", Title: "Go Programming", Author: "Gopher"})
	s.AddBook(models.Book{ID: "b2", Title: "Rust Essentials", Author: "Ferris"})

	results := s.SearchBooks("go")
	if len(results) != 1 || results[0].ID != "b1" {
		t.Fatalf("expected to find only Go book, got: %+v", results)
	}

	results = s.SearchBooks("")
	if len(results) != 2 {
		t.Fatalf("empty search should return all books")
	}
}

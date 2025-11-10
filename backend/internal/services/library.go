package services

import (
	"errors"
	"strings"

	"library/internal/ds"
	"library/internal/models"
)

type LibraryService struct {
	books       *ds.BST[string, models.Book]
	users       *ds.BST[string, models.User]
	activeLoans *ds.BST[string, models.LoanRequest]
	history     *ds.Stack[string]
	featured    *ds.Array[string]
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		books:       ds.NewBST[string, models.Book](strings.Compare),
		users:       ds.NewBST[string, models.User](strings.Compare),
		activeLoans: ds.NewBST[string, models.LoanRequest](strings.Compare),
		history:     ds.NewStack[string](),
		featured:    ds.NewArray[string](5),
	}
}

func (s *LibraryService) AddBook(b models.Book) {
	b.Available = true
	s.books.Put(b.ID, b)
	s.history.Push("add_book:" + b.ID)
}

func (s *LibraryService) ListBooks() []models.Book {
	out := make([]models.Book, 0, s.books.Size())
	s.books.TraverseInOrder(func(_ string, v models.Book) { out = append(out, v) })
	return out
}

func (s *LibraryService) SearchBooks(q string) []models.Book {
	q = strings.ToLower(strings.TrimSpace(q))
	out := make([]models.Book, 0)
	if q == "" {
		return s.ListBooks()
	}
	s.books.TraverseInOrder(func(_ string, v models.Book) {
		if strings.Contains(strings.ToLower(v.Title), q) || strings.Contains(strings.ToLower(v.Author), q) {
			out = append(out, v)
		}
	})
	return out
}

func (s *LibraryService) AddUser(u models.User) {
	s.users.Put(u.ID, u)
	s.history.Push("add_user:" + u.ID)
}

func (s *LibraryService) ListUsers() []models.User {
	out := make([]models.User, 0, s.users.Size())
	s.users.TraverseInOrder(func(_ string, v models.User) { out = append(out, v) })
	return out
}

func (s *LibraryService) Borrow(req models.LoanRequest) error {
	if _, ok := s.users.Get(req.UserID); !ok {
		return errors.New("user not found")
	}
	book, ok := s.books.Get(req.BookID)
	if !ok {
		return errors.New("book not found")
	}
	if !book.Available {
		return errors.New("book not available")
	}
	if _, exists := s.activeLoans.Get(req.BookID); exists {
		return errors.New("book already loaned")
	}
	book.Available = false
	s.books.Put(book.ID, book)
	s.activeLoans.Put(req.BookID, req)
	s.history.Push("borrow:" + req.UserID + ":" + req.BookID)
	return nil
}

func (s *LibraryService) Return(req models.LoanRequest) error {
	loan, ok := s.activeLoans.Get(req.BookID)
	if !ok {
		return errors.New("loan not found")
	}
	if loan.UserID != req.UserID {
		return errors.New("loan belongs to a different user")
	}
	book, ok := s.books.Get(req.BookID)
	if !ok {
		return errors.New("book not found")
	}
	book.Available = true
	s.books.Put(book.ID, book)
	s.activeLoans.Delete(req.BookID)
	s.history.Push("return:" + req.UserID + ":" + req.BookID)
	return nil
}

func (s *LibraryService) HistorySize() int { return s.history.Size() }

// RemoveBook deletes a book by ID. It refuses to delete if the book is currently loaned (Available=false).
func (s *LibraryService) RemoveBook(id string) error {
	if _, active := s.activeLoans.Get(id); active {
		return errors.New("book currently loaned")
	}
	_, deleted := s.books.Delete(id)
	if !deleted {
		return errors.New("book not found")
	}
	s.history.Push("remove_book:" + id)
	return nil
}

// RemoveUser deletes a user by ID.
func (s *LibraryService) RemoveUser(id string) error {
	hasLoans := false
	s.activeLoans.TraverseInOrder(func(_ string, loan models.LoanRequest) {
		if loan.UserID == id {
			hasLoans = true
		}
	})
	if hasLoans {
		return errors.New("user has active loans")
	}
	_, deleted := s.users.Delete(id)
	if !deleted {
		return errors.New("user not found")
	}
	s.history.Push("remove_user:" + id)
	return nil
}

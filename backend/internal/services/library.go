package services

import (
	"errors"
	"strings"

	"library/internal/ds"
	"library/internal/models"
)

type LibraryService struct {
	books    *ds.List[models.Book]
	users    *ds.List[models.User]
	loans    *ds.Queue[models.LoanRequest]
	history  *ds.Stack[string]
	featured *ds.Array[string]
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		books:    ds.NewList[models.Book](),
		users:    ds.NewList[models.User](),
		loans:    ds.NewQueue[models.LoanRequest](),
		history:  ds.NewStack[string](),
		featured: ds.NewArray[string](5),
	}
}

func (s *LibraryService) AddBook(b models.Book) {
    b.Available = true
    s.books.InsertFront(b)
    s.history.Push("add_book:" + b.ID)
}

func (s *LibraryService) ListBooks() []models.Book {
    // Initialize with empty slice to avoid null in JSON
    out := make([]models.Book, 0)
    s.books.ForEach(func(v models.Book) { out = append(out, v) })
    return out
}

func (s *LibraryService) SearchBooks(q string) []models.Book {
    q = strings.ToLower(strings.TrimSpace(q))
    out := make([]models.Book, 0)
    if q == "" {
        return s.ListBooks()
    }
    s.books.ForEach(func(v models.Book) {
        if strings.Contains(strings.ToLower(v.Title), q) || strings.Contains(strings.ToLower(v.Author), q) {
            out = append(out, v)
        }
    })
    return out
}

func (s *LibraryService) AddUser(u models.User) {
	s.users.InsertFront(u)
	s.history.Push("add_user:" + u.ID)
}

func (s *LibraryService) ListUsers() []models.User {
    out := make([]models.User, 0)
    s.users.ForEach(func(v models.User) { out = append(out, v) })
    return out
}

func (s *LibraryService) Borrow(req models.LoanRequest) error {
    s.loans.Enqueue(req)
    return s.processNextLoan()
}

func (s *LibraryService) processNextLoan() error {
	req, ok := s.loans.Dequeue()
	if !ok {
		return nil
	}
	// find book
	found := false
	// rebuild list with updated availability
	var tmp []models.Book
	s.books.ForEach(func(v models.Book) { tmp = append(tmp, v) })
	for i := range tmp {
		if tmp[i].ID == req.BookID {
			found = true
			if !tmp[i].Available {
				return errors.New("book not available")
			}
			tmp[i].Available = false
			break
		}
	}
	if !found {
		return errors.New("book not found")
	}
	// rebuild list: clear and reinsert
	s.books = ds.NewList[models.Book]()
	for i := len(tmp) - 1; i >= 0; i-- {
		s.books.InsertFront(tmp[i])
	}
	s.history.Push("borrow:" + req.UserID + ":" + req.BookID)
	return nil
}

func (s *LibraryService) Return(req models.LoanRequest) error {
	var tmp []models.Book
	s.books.ForEach(func(v models.Book) { tmp = append(tmp, v) })
	found := false
	for i := range tmp {
		if tmp[i].ID == req.BookID {
			tmp[i].Available = true
			found = true
			break
		}
	}
	if !found {
		return errors.New("book not found")
	}
	s.books = ds.NewList[models.Book]()
	for i := len(tmp) - 1; i >= 0; i-- {
		s.books.InsertFront(tmp[i])
	}
	s.history.Push("return:" + req.UserID + ":" + req.BookID)
	return nil
}

func (s *LibraryService) HistorySize() int { return s.history.Size() }

// RemoveBook deletes a book by ID. It refuses to delete if the book is currently loaned (Available=false).
func (s *LibraryService) RemoveBook(id string) error {
    var tmp []models.Book
    s.books.ForEach(func(v models.Book) { tmp = append(tmp, v) })
    found := false
    for i := range tmp {
        if tmp[i].ID == id {
            if !tmp[i].Available {
                return errors.New("book currently loaned")
            }
            // remove i
            tmp = append(tmp[:i], tmp[i+1:]...)
            found = true
            break
        }
    }
    if !found { return errors.New("book not found") }
    s.books = ds.NewList[models.Book]()
    for i := len(tmp) - 1; i >= 0; i-- { s.books.InsertFront(tmp[i]) }
    s.history.Push("remove_book:" + id)
    return nil
}

// RemoveUser deletes a user by ID.
func (s *LibraryService) RemoveUser(id string) error {
    var tmp []models.User
    s.users.ForEach(func(v models.User) { tmp = append(tmp, v) })
    found := false
    for i := range tmp {
        if tmp[i].ID == id {
            tmp = append(tmp[:i], tmp[i+1:]...)
            found = true
            break
        }
    }
    if !found { return errors.New("user not found") }
    s.users = ds.NewList[models.User]()
    for i := len(tmp) - 1; i >= 0; i-- { s.users.InsertFront(tmp[i]) }
    s.history.Push("remove_user:" + id)
    return nil
}

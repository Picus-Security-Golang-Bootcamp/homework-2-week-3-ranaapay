package service

import (
	"PicusHomework2/src/storage"
	"PicusHomework2/src/storage/storageType"
	"fmt"
)

type Service struct {
	library *storage.Library
}

func NewService(library *storage.Library) *Service {
	return &Service{library: library}
}

func (s *Service) WriteBookList() {
	bookList := s.library.FindBooksNotDeleted()
	s.printBookList(bookList)
}

func (s *Service) SearchBooks(word string) {
	books := s.library.FindBooksByItsName(word)
	if books == nil {
		fmt.Printf("Book with %s name is not contain in the book list.", word)
	} else {
		s.printBookList(books)
	}
}

func (s *Service) WriteBookById(id int) {
	book := s.library.FindBookById(id)
	if book == nil {
		fmt.Printf("Book with %d id is not contain in the book list.", id)
	} else {
		book.PrintBook()
	}
}

func (s *Service) DeleteBookById(id int) {
	isDeleted := s.library.DeleteBookById(id)
	if !isDeleted {
		fmt.Printf("Book with %d id is not contain in the book list.", id)
	} else {
		fmt.Printf("The book with id : %d has been deleted from the book list.", id)
	}
}

func (s *Service) BuyBookByItsId(id, quantity int) {
	book, buyNum := s.library.BuyBookByItsId(id, quantity)
	if book == nil {
		fmt.Printf("Can not buy any books. Please confirm parameters.")
	} else {
		book.PrintBook()
		fmt.Printf("Buy : %d", buyNum)
	}
}

func (s *Service) SaveChanges() {
	s.library.SaveChanges()
}

func (s *Service) printBookList(bookList []storageType.Book) {
	for _, book := range bookList {
		book.PrintBook()
	}
}

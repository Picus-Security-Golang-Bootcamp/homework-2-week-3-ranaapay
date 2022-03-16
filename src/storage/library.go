package storage

import (
	"PicusHomework2/src/storage/storageType"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Library struct {
	bookList []storageType.Book
}

func NewLibrary() *Library {
	library := Library{}
	library.insertBooks()
	return &library
}

//insertBooks Reads books from bookData.json file and assign it to BookList in Library.
func (l *Library) insertBooks() {
	var bookList []storageType.Book
	jsonFile, err := os.Open("bookData.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err = json.Unmarshal(byteValue, &bookList); err != nil {
		fmt.Println("Somethings went wrong when unmarshal data.")
	}
	l.bookList = bookList
}

//FindBooksNotDeleted find books that deleted field is not true and return list
func (l *Library) FindBooksNotDeleted() []storageType.Book {
	var bookList []storageType.Book
	for _, book := range l.bookList {
		if book.GetIsDeleted() == false {
			bookList = append(bookList, book)
		}
	}
	return bookList
}

//FindBooksByItsName check if book name contains the word.
//if it contains append to bookList and return list.
func (l *Library) FindBooksByItsName(searchWord string) []storageType.Book {
	var books []storageType.Book
	searchWord = strings.ToLower(searchWord)
	for _, book := range l.bookList {
		bLow := strings.ToLower(book.GetBookName())
		index := strings.Index(bLow, searchWord)
		if index == -1 {
			continue
		}
		books = append(books, book)
	}
	return books
}

//FindBookById find and return the book that Book.Id is equal to id value
func (l *Library) FindBookById(id int) *storageType.Book {
	for _, book := range l.bookList {
		if book.GetBookId() == id {
			return &book
		}
	}
	return nil
}

//DeleteBookById changes the book isDeleted field to true by its id
func (l *Library) DeleteBookById(id int) bool {
	bookIndex := -1
	for i, book := range l.bookList {
		if book.GetBookId() == id {
			bookIndex = i
		}
	}
	if bookIndex >= 0 {
		l.bookList[bookIndex].SetIsDeleted(true)
		return true
	}
	return false
}

//BuyBookByItsId checks the stock number of the book according to the entered id value
//and quantity and returns the number of books bought and the latest status of the book.
func (l *Library) BuyBookByItsId(id, quantity int) (*storageType.Book, int) {
	bookIndex := -1
	for i, book := range l.bookList {
		if book.GetBookId() == id {
			bookIndex = i
		}
	}
	if bookIndex >= 0 {
		b := l.bookList[bookIndex].GetStockNumber()
		if b == 0 {
			return nil, 0
		}
		if b < quantity {
			l.bookList[bookIndex].SetStockNumber(0)
			return &l.bookList[bookIndex], b
		} else {
			newStock := b - quantity
			l.bookList[bookIndex].SetStockNumber(newStock)
			return &l.bookList[bookIndex], quantity
		}
	}
	return nil, 0
}

//SaveChanges writes the final version of the book list to the file
func (l *Library) SaveChanges() {
	content, err := json.Marshal(l.bookList)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = ioutil.WriteFile("bookData.json", content, 0644); err != nil {
		fmt.Println(err.Error())
	}
}

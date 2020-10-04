package repository

import (
	. "github.com/hsedjame/restapi-mux/src/models"
	"math/rand"
	"strconv"
)

type BookRepository interface {
	
	GetBooks() []Book
	
	GetBook(id string) Book
	
	CreateBook(book Book) []Book
	
	UpdateBook(book Book) []Book
	
	DeleteBook(id string) bool
}

type BookRepositoryImpl struct{
	Books []Book
}

func NewBookRepositoryImpl() *BookRepositoryImpl {
	return &BookRepositoryImpl{Books: []Book{}}
}

func (repo BookRepositoryImpl) GetBooks() []Book {
	return repo.Books
}

func (repo BookRepositoryImpl) GetBook(id string) Book {
	var book Book

	for _,b := range repo.Books {
		if b.ID == id {
			book = b
			break
		}
	}
	return book
}

func (repo *BookRepositoryImpl) CreateBook(book Book) []Book {

	i := rand.Intn(100000000)
	id := strconv.Itoa(i)
	book.ID = id
	repo.Books = append(repo.Books, book)
	return repo.Books
}

func (repo *BookRepositoryImpl) UpdateBook(book Book) []Book {

	for i, b := range repo.Books {
		if b.ID == book.ID {
			repo.Books[i] = book
			break
		}
	}

	return repo.Books
}

func (repo *BookRepositoryImpl) DeleteBook(id string) bool {
	for i, b := range repo.Books {
		if b.ID == id {
			repo.Books = append(repo.Books[:i], repo.Books[(i+1):]...)
			break
		}
	}
	return true
}


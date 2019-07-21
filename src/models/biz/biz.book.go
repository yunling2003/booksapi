package biz

import (
	"booksapi/src/models/orm"
)

func GetAllBooks(books *[]orm.Book) error {
	var bookOrm orm.Book
	return bookOrm.GetAll(books)
}
package data

import (
	"api/features/book"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.BookData {
	return &bookData{
		db: db,
	}
}

func (bd *bookData) Add(userID int, newBook book.Core) (book.Core, error) {
	cnv := CoreToData(newBook)
	cnv.UserID = uint(userID)
	err := bd.db.Create(&cnv).Error
	if err != nil {
		return book.Core{}, err
	}

	newBook.ID = cnv.ID

	return newBook, nil
}
func (bd *bookData) Update(bookID, updatedData book.Core) (book.Core, error) {
	return book.Core{}, nil
}
func (bd *bookData) Delete(bookID int) error {
	return nil
}
func (bd *bookData) MyBook(userID int) ([]book.Core, error) {
	return nil, nil
}

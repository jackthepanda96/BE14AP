package book

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Judul       string
	TahunTerbit int
	Penulis     string
	Pemilik     string
}

type BookHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyBook() echo.HandlerFunc
}

type BookService interface {
	Add(newBook Core) (Core, error)
	Update(bookID, updatedData Core) (Core, error)
	Delete(bookID int) error
	MyBook(userID int) ([]Core, error)
}

type BookData interface {
	Add(userID int, newBook Core) (Core, error)
	Update(bookID, updatedData Core) (Core, error)
	Delete(bookID int) error
	MyBook(userID int) ([]Core, error)
}

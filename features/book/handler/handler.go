package handler

import (
	"api/features/book"
	"api/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookHandle struct {
	srv book.BookService
}

func New(bs book.BookService) book.BookHandler {
	return &bookHandle{
		srv: bs,
	}
}

func (bh *bookHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ToCore(input)

		res, err := bh.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan buku", res))
	}
}
func (bh *bookHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusNoContent, "")
	}
}

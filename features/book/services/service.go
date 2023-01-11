package services

import (
	"api/features/book"
	"api/helper"
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type bookSrv struct {
	data book.BookData
	vld  *validator.Validate
}

func New(d book.BookData) book.BookService {
	return &bookSrv{
		data: d,
		vld:  validator.New(),
	}
}

func (bs *bookSrv) Add(token interface{}, newBook book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return book.Core{}, errors.New("user tidak ditemukan")
	}

	err := bs.vld.Struct(newBook)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return book.Core{}, errors.New("input buku tidak sesuai dengan arahan")
	}

	res, err := bs.data.Add(userID, newBook)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "buku tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return book.Core{}, errors.New(msg)
	}

	return res, nil

}
func (bs *bookSrv) Update(token interface{}, bookID int, updatedData book.Core) (book.Core, error) {
	return book.Core{}, nil
}

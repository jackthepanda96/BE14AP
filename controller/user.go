package controller

import (
	"api/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Tugas controller, membaca input dari user
// meneruskan input ke bagian proses,
// memberikan respon ke user
type UserControll struct {
	Mdl model.UserModel
}

func (uc *UserControll) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp := model.User{}
		err := c.Bind(&tmp)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.Mdl.Insert(tmp)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "sukses menambahkan data"})
	}
}

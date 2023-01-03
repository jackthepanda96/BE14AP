package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Name     string `json:"name" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Email    string `json:"email" form:"email"`
	Password string `json:"pwd" form:"password"`
}

type UserControll struct {
	DB *gorm.DB
}

func (uc *UserControll) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp := User{}
		err := c.Bind(&tmp)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		err = uc.DB.Create(tmp).Error
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusCreated, "user berahsil di tambahkan")
	}
}

func (uc *UserControll) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		res := []User{}

		// err := uc.DB.Find(&res).Error // get all -> select *

		err := uc.DB.Where("name = ?", "kamil").Find(&res).Error

		if err != nil {
			log.Println("query error", err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "berhasil menampilkan daftar user",
		})
	}
}

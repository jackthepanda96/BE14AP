package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Nama     string `json:"name" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Email    string `json:"email" form:"email"`
	Password string `json:"pwd" form:"password"`
}

type UserControll struct {
	ListUser []User
}

func (uc *UserControll) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp := User{}
		err := c.Bind(&tmp)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		uc.ListUser = append(uc.ListUser, tmp)

		return c.JSON(http.StatusCreated, "user berahsil di tambahkan")
	}
}

func (uc *UserControll) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    uc.ListUser,
			"message": "berhasil menampilkan daftar user",
		})
	}
}

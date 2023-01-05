package handler

import (
	"api/features/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil login", res))
	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", res))
	}
}
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := c.Param("id")
		cnv, err := strconv.Atoi(input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "id salah")
		}

		res, err := uc.srv.Profile(uint(cnv))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", res))
	}
}

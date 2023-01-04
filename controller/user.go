package controller

import (
	"api/model"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// Tugas controller, membaca input dari user
// meneruskan input ke bagian proses,
// memberikan respon ke user
type UserControll struct {
	Mdl model.UserModel
}

func (uc *UserControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp := model.User{}
		if err := c.Bind(&tmp); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.Mdl.Login(tmp.Email, tmp.Password)

		if err != nil {
			if strings.Contains(err.Error(), "matched") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": err.Error(),
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "sukses login",
		})
	}
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

func (uc *UserControll) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.Mdl.GetAll()
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "sukses mendapatkan semua data"})
	}
}

// /users/:id
func (uc *UserControll) GetID() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		cnvID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gunakan input angka",
			})
		}
		res, err := uc.Mdl.GetByID(cnvID)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "sukses mendapatkan semua data"})
	}
}

func (uc *UserControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		cnvID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gunakan input angka",
			})
		}
		body := model.User{}
		err = c.Bind(&body)
		if err != nil {
			log.Println("bind body error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "masukkan input sesuai pola",
			})
		}
		body.ID = uint(cnvID)
		res, err := uc.Mdl.Update(body)

		if err != nil {
			log.Println("query error ", err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "berhasil update data",
		})

	}
}

func (uc *UserControll) Update2() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		cnvID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gunakan input angka",
			})
		}
		body := model.User{}
		err = c.Bind(&body)
		if err != nil {
			log.Println("bind body error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "masukkan input sesuai pola",
			})
		}
		body.ID = uint(cnvID)
		res, err := uc.Mdl.Update2(body)

		if err != nil {
			log.Println("query error ", err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "berhasil update data",
		})
	}
}

func (uc *UserControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		cnvID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gunakan input angka",
			})
		}

		err = uc.Mdl.Delete(cnvID)

		if err != nil {
			log.Println("delete error", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sukses menghapus data",
		})
	}
}

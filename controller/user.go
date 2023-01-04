package controller

import (
	"api/model"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Tugas controller, membaca input dari user
// meneruskan input ke bagian proses,
// memberikan respon ke user
type UserControll struct {
	Mdl model.UserModel
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("BE!4a|t3rr4"))
}

func ExtractToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(float64)
		return int(userId)
	}
	return -1
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
		jwtToken, err := CreateToken(int(res.ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"token":   jwtToken,
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
		id := ExtractToken(c)
		// paramID := c.Param("id")
		// cnvID, err := strconv.Atoi(paramID)
		// if err != nil {
		// 	log.Println("convert id error ", err.Error())
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"message": "gunakan input angka",
		// 	})
		// }
		res, err := uc.Mdl.GetByID(id)
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
		id := ExtractToken(c)
		// paramID := c.Param("id")
		// cnvID, err := strconv.Atoi(paramID)
		// if err != nil {
		// 	log.Println("convert id error ", err.Error())
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"message": "gunakan input angka",
		// 	})
		// }
		body := model.User{}
		err := c.Bind(&body)
		if err != nil {
			log.Println("bind body error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "masukkan input sesuai pola",
			})
		}
		body.ID = uint(id)
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
		id := ExtractToken(c)
		// paramID := c.Param("id")
		// cnvID, err := strconv.Atoi(paramID)
		// if err != nil {
		// 	log.Println("convert id error ", err.Error())
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"message": "gunakan input angka",
		// 	})
		// }
		body := model.User{}
		err := c.Bind(&body)
		if err != nil {
			log.Println("bind body error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "masukkan input sesuai pola",
			})
		}
		body.ID = uint(id)
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
		id := ExtractToken(c)
		// paramID := c.Param("id")
		// cnvID, err := strconv.Atoi(paramID)
		// if err != nil {
		// 	log.Println("convert id error ", err.Error())
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"message": "gunakan input angka",
		// 	})
		// }

		err := uc.Mdl.Delete(id)

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

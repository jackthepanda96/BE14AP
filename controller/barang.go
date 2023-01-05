package controller

import (
	"api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GoodsControll struct {
	Mdl model.GoodsModel
}

func (gc *GoodsControll) AddGoods() echo.HandlerFunc {
	return func(c echo.Context) error {
		// extract token
		id := ExtractToken(c)
		if id <= 0 {
			return c.JSON(http.StatusUnauthorized, "tidak boleh mengakses kesini")
		}

		input := model.Goods{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		input.UserID = uint(id)

		res, err := gc.Mdl.Insert(input)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error dari server")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "berhasil menambahkan barang",
		})

	}
}

func (gc *GoodsControll) ShowGoods() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := ExtractToken(c)
		if id <= 0 {
			return c.JSON(http.StatusUnauthorized, "tidak boleh mengakses kesini")
		}

		res, err := gc.Mdl.GetAllByID(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error dari server")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "berhasil menambahkan barang",
		})
	}
}

func (gc *GoodsControll) GoodsDetails() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := ExtractToken(c)
		if id <= 0 {
			return c.JSON(http.StatusUnauthorized, "tidak boleh mengakses kesini")
		}

		res, err := gc.Mdl.Details(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error dari server")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "lihat data barang",
		})
	}
}

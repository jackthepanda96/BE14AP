package main

import (
	"api/user"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// fmt.Println("Run server....")
//
//	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == "GET" {
//			w.Header().Set("content-type", "application/json")
//			w.Write([]byte("Hello world"))
//		}
//	})
//
// err := http.ListenAndServe(":8000", nil)
//
//	if err != nil {
//		log.Println("Start server error")
//	}
func main() {
	e := echo.New()
	arr := []user.User{}
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World")
	})
	e.POST("/users", func(c echo.Context) error {
		tmp := user.User{}
		err := c.Bind(&tmp)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		arr = append(arr, tmp)

		return c.JSON(http.StatusCreated, "user berahsil di tambahkan")
	})
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    arr,
			"message": "berhasil menampilkan daftar user",
		})
	})
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

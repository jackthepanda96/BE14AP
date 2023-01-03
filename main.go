package main

import (
	"api/config"
	"api/user"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)
	controll := user.UserControll{DB: db}

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World")
	})
	e.POST("/users", controll.Insert())
	e.GET("/users", controll.GetAllUser())
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

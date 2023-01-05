package main

import (
	"api/config"
	"api/features/user/data"
	"api/features/user/handler"
	"api/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users/:id", userHdl.Profile())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

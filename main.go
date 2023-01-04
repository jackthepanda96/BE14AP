package main

import (
	"api/config"
	"api/controller"
	"api/model"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)
	model := model.UserModel{DB: db}
	controll := controller.UserControll{Mdl: model}

	e.POST("/users", controll.Insert())
	e.POST("/login", controll.Login())
	e.GET("/users", controll.GetAll())
	e.GET("/users/:id", controll.GetID())
	e.PATCH("/users/:id", controll.Update())
	e.PUT("/users/:id", controll.Update2())
	e.DELETE("/users/:id", controll.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

package main

import (
	"api/config"
	"api/controller"
	"api/model"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)
	model := model.UserModel{DB: db}
	controll := controller.UserControll{Mdl: model}

	e.Pre(middleware.RemoveTrailingSlash()) // fungsi ini dijalankan sebelum routing

	e.Use(middleware.CORS()) // WAJIB DIPAKAI agar tidak terjadi masalah permission
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.Use(middleware.Logger()) // Dipakai untuk membuat log (catatan) ketika endpoint diakses

	e.POST("/register", controll.Insert())

	e.POST("/login", controll.Login())

	needLogin := e.Group("/users")
	// needLogin.Use(middleware.JWT([]byte(cfg.JWTKey)))
	needLogin.GET("", controll.GetAll())
	needLogin.GET("/:id", controll.GetID(), middleware.JWT([]byte(config.JWT_KEY)))
	needLogin.PATCH("", controll.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	// PATCH localhost:8000/users/:id/patch
	needLogin.PUT("", controll.Update2(), middleware.JWT([]byte(config.JWT_KEY)))
	needLogin.DELETE("", controll.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}

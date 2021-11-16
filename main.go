package main

import (
	"ytmy-bg/controller"
	"ytmy-bg/lib"
	"ytmy-bg/middle"
	"ytmy-bg/model"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.POST("/api/admin/login", controller.AdminLoginPost)
	login := r.Group("/api")
	login.Use(middle.IsLogin())
	{
		login.GET("/goods", controller.GoodGet)
		admin := login.Group("")
		admin.Use(middle.IsAdmin())
		{
			admin.POST("/goods", controller.GoodPost)
			admin.PUT("/goods/:id", controller.GoodPut)
			admin.DELETE("goods/:id", controller.GoodDelete)
		}
	}

	return r
}

func dbinstance() {
	db := lib.GetDB()
	db.AutoMigrate(
		&model.Admin{},
		&model.Good{},
	)
}

func main() {
	r := setupRouter()
	dbinstance()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

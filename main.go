package main

import (
	"ytmy-bg/controller"
	"ytmy-bg/lib"
	"ytmy-bg/middle"
	"ytmy-bg/model"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.POST("/api/admin/login", controller.AdminLoginPost)
	r.POST("/api/user/login", controller.UserLogin)
	login := r.Group("/api")
	login.Use(middle.IsLogin())
	{
		login.GET("/goods", controller.GoodGet)
		login.GET("/carts", controller.CartGet)
		login.POST("/carts", controller.CartPost)
		login.PUT("/carts/:id", controller.CartPut)
		login.DELETE("/carts/:id", controller.CartDelete)
		login.GET("/orders", controller.OrderGet)
		login.POST("/orders", controller.OrderPost)
		admin := login.Group("")
		admin.Use(middle.IsAdmin())
		{
			admin.POST("/goods", controller.GoodPost)
			admin.PUT("/goods/:id", controller.GoodPut)
			admin.DELETE("/goods/:id", controller.GoodDelete)
			admin.PUT("/orders/:id", controller.OrderPut)
		}
	}

	return r
}

func dbinstance() {
	db := lib.DB()
	db.AutoMigrate(
		&model.Admin{},
		&model.Good{},
		&model.Order{},
		&model.Cart{},
		&model.User{},
	)
}

func main() {
	r := setupRouter()
	dbinstance()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

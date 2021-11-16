package main

import (
	"ytmy-bg/controller"
	"ytmy-bg/lib"
	"ytmy-bg/model"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.POST("/api/admin/login", controller.AdminLoginPost)
	return r
}

func dbinstance() {
	db := lib.GetDB()
	db.AutoMigrate(
		&model.Admin{},
	)
}

func main() {
	r := setupRouter()
	dbinstance()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

package controller

import (
	"net/http"
	"ytmy-bg/lib"
	"ytmy-bg/model"

	"github.com/gin-gonic/gin"
)

func CartPost(c *gin.Context) {
	cart := model.Cart{}
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	result := db.Create(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func CartGet(c *gin.Context) {
	var Carts []model.Cart
	db := lib.DB()
	userid := c.GetInt("userID")
	result := db.Where("userid = ?", userid).Find(&Carts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	} else {
		c.JSON(http.StatusOK, Carts)
	}
}

func CartPut(c *gin.Context) {
	Cart := model.Cart{}
	if err := c.ShouldBindJSON(&Cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var CartDB model.Cart
	db.First(&CartDB, id)
	CartDB.Num = Cart.Num
	db.Save(&CartDB)
	c.JSON(http.StatusOK, "OK")
}

func CartDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	db.Delete(&model.Cart{}, id)
	c.JSON(http.StatusOK, "OK")
}

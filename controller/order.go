package controller

import (
	"net/http"
	"ytmy-bg/lib"
	"ytmy-bg/model"

	"github.com/gin-gonic/gin"
)

func OrderPost(c *gin.Context) {
	order := model.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := lib.DB()
	userID := c.GetUint("userID")
	order.UserID = userID
	order.Status = 1
	result := db.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func OrderGet(c *gin.Context) {
	var Orders []model.Order
	db := lib.DB()
	userid := c.GetUint("userID")
	result := db.Where("user_id = ?", userid).Preload("Carts").Order("status").Find(&Orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
	} else {
		c.JSON(http.StatusOK, Orders)
	}
}

func OrderPut(c *gin.Context) {
	Order := model.Order{}
	if err := c.ShouldBindJSON(&Order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	db := lib.DB()
	var OrderDB model.Order
	userID := c.GetInt("userID")
	result := db.Where("user_id = ?", userID).First(&OrderDB, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	OrderDB.Status = Order.Status
	result = db.Save(&OrderDB)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func OrderDelete(c *gin.Context) {
	id := c.Param("id")
	db := lib.DB()
	userID := c.GetInt("userID")
	result := db.Where("user_id = ?", userID).Delete(&model.Order{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

package controller

import (
	"net/http"
	"ytmy-bg/lib"
	"ytmy-bg/model"

	"github.com/gin-gonic/gin"
)

func AdminLoginPost(c *gin.Context) {
	admin := model.Admin{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	db := lib.GetDB()
	adminDB := model.Admin{}
	db.Where("name=?", admin.Name).First(&adminDB)
	if adminDB.Pwd == lib.MD5(admin.Pwd) {
		cookie := lib.Cookie{ID: adminDB.ID, Type: 1}
		cookieStr, err := lib.CookieEncrypt(cookie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.SetCookie("ytmy", cookieStr, 3600, "/", "*", false, true)
			c.JSON(http.StatusOK, "OK")
		}
	} else {
		c.JSON(http.StatusBadRequest, "用户名或密码错误")
	}
}

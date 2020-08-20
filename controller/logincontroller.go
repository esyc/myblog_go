package controller

import (
	"crcblog/middleware"
	"crcblog/models"
	"crcblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
}

func UserLogin(ctx *gin.Context)  {
	var data models.User
	var token string
	var code int
	ctx.ShouldBindJSON(&data)
	code = models.LoginCheck(data.UserName,data.UserPassword)

	if code ==utils.SUCCESS{
		token,code = middleware.SetToken(data.UserName,data.UserName)
	}
		ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"token":token,
		"message":utils.ErrorMsg(code),
	})

}

package controller

import (
	"crcblog/models"
	"crcblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserExist(cxt *gin.Context)  {

}

func AddUser(ctx *gin.Context){
	var data models.User
	_ = ctx.ShouldBindJSON(&data)
	code := models.UserCheck(data.UserName)
	if code != utils.SUCCESS{
		code = models.InsertUser(&data)
	} else if code == utils.SUCCESS{
		code = utils.REPEAT_NAME
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})

}
func DeleteUser(ctx *gin.Context){
	UserName  := ctx.Query("name")
	code := models.UserCheck(UserName)
	if code == utils.SUCCESS{
		code = models.DeleteUser(UserName)
	}
	if code == utils.ERROR {
		code = utils.USER_NOT_EXIST
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"name":UserName,
		"message":utils.ErrorMsg(code),
	})
}
func GetUser(cxt *gin.Context)  {
	
}

func EditUser(cxt *gin.Context){

}
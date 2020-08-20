package controller

import (
	"crcblog/models"
	"crcblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage (ctx *gin.Context){
	data,code:= models.GetMessage()
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
func AddMessage (ctx *gin.Context){
	var data models.Message
	_ = ctx.ShouldBindJSON(&data)
	models.InsertMessage(data)
	code := utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":utils.ErrorMsg(code),
	})
}
func GetZan (ctx *gin.Context){
	data,code:= models.GetZan()
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
func AddZan (ctx *gin.Context){
	models.UpdateZan()
	code := utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":utils.ErrorMsg(code),
	})
}
func GetPeo (ctx *gin.Context){
	data,code:= models.GetPeo()
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
func AddPeo (ctx *gin.Context){
	models.UpdatePeo()
	code := utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":utils.ErrorMsg(code),
	})
}
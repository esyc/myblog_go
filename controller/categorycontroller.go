package controller

import (
	"crcblog/models"
	"crcblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCategory(ctx *gin.Context){
	data:=models.GetAllCate()
	code := utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}

func DeleteCategory(ctx *gin.Context){
	id,_:= strconv.Atoi(ctx.Param("id"))
	code:= models.DeleteCate(id)
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":utils.ErrorMsg(code),
	})
}
func AddCategory(ctx *gin.Context){
	var data models.Category
	_ = ctx.ShouldBindJSON(&data)
	code := models.CateCheck(data.CateName)
	if code != utils.SUCCESS{
		code = models.InsertCate(&data)
	} else if code == utils.SUCCESS{
		code = utils.REPEAT_CATE
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
func EditCategory(ctx *gin.Context)  {
	var data models.Category
	id,_:=strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)
	code := models.EditCate(id,&data)
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":utils.ErrorMsg(code),
	})
}

package controller

import (
	"crcblog/models"
	"crcblog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)
//添加文章
func AddPost(ctx *gin.Context)  {
	var data models.Post
	_ = ctx.ShouldBindJSON(&data)

	data.PostTime = time.Now()
	code := models.InsertPost(&data)
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
//查询分类下文章
func GetPostFromCate(ctx *gin.Context){
	cid,_:=strconv.Atoi(ctx.Param("cid"))
	data := models.GetPostFromCate(cid)
	code :=utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
//查询所有文章
func GetAllPost(ctx *gin.Context)  {
	data:=models.GetAllPost()
	code := utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
//分页查询文章
func GetPostByNum(ctx *gin.Context) {
	pageSize,_ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum,_ := strconv.Atoi(ctx.Query("pagenum"))
	data:=models.GetPostByNum(pageSize,pageNum)
	code := utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})

}
//查询单个文章
func GetOnePost(ctx *gin.Context)  {
	id,_:=strconv.Atoi(ctx.Param("id"))
	data,code:= models.GetOnePost(id)
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":utils.ErrorMsg(code),
	})
}
//修改文章
func EditPost(ctx *gin.Context)  {
	var data models.Post
	id,_:=strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)
	code := models.EditPost(id,&data)
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":utils.ErrorMsg(code),
	})
}
//删除文章
func DeletePost(ctx *gin.Context){
	id,_:= strconv.Atoi(ctx.Param("id"))
	code:= models.DeletePost(id)
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":utils.ErrorMsg(code),
	})
}
func GetPostNum(ctx *gin.Context){
	n := models.GetPostNum()
	data :=  strconv.FormatInt(n,10)
	fmt.Println(data)
	code:=utils.SUCCESS
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"num": data,
		"message":utils.ErrorMsg(code),
	})
}
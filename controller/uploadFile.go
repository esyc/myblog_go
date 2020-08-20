package controller

import (
	"crcblog/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadFile (ctx *gin.Context){
	file,err :=ctx.FormFile("images")
	if err !=nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":500,
			"message": utils.ErrorMsg(500),
		})
		return
	}
	rand.Seed(time.Now().Unix())
	rnum :=rand.Int63n(1000000000)
	rnum = rnum + time.Now().Unix()
	filename :=  strconv.FormatInt(rnum,10)
	filename = filename + file.Filename
	if err := ctx.SaveUploadedFile(file,"./static/img/"+filename); err != nil{
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"message":utils.ErrorMsg(500),
		})
		return
	}
	filename = "http://localhost:8889/file/img/"+filename
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":filename,
		"message":utils.ErrorMsg(200),
	})
}
func GetFile (ctx *gin.Context){
	filename:= ctx.Param("filename")
	path := "./static/img/"
	path += filename
	ctx.File(path)
}
func DeleteFile (ctx *gin.Context){
	filename:= ctx.Param("filename")
	path := "./static/img/"
	path += filename
	err := os.Remove(path)
	if err!=nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    filename,
			"message": utils.ErrorMsg(500),
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":filename,
		"message":utils.ErrorMsg(200),
	})

}
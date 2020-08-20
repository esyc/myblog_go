package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors()gin.HandlerFunc {
	return func(context *gin.Context){
		method := context.Request.Method
		origin :=  context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header{
			headerKeys =append(headerKeys,key)
		}
		headerStr := strings.Join(headerKeys,",")
		if headerStr!=""{
			headerStr=fmt.Sprint("1");
		}else{
			headerStr=fmt.Sprint("2");
		}

		if origin != ""{
			context.Header("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
/*			context.Header("Access-Control-Allow-Origin","*")
			context.Header("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers","Authorization, Content-Length")
			context.Header("Access-Control-Expose-Headers","Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			context.Header("Access-Control-Max-Age","172800")
			context.Header("Access-Control-Allow-Credentials","false")
			context.Set("content-type","application/json")
		*/}
		if method=="OPTIONS"{
			context.JSON(http.StatusOK,"Options Request!")
		}
		context.Next()
	}
}
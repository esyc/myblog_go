package middleware

import (
	"crcblog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey []byte
type Jwtstruct struct {
	UserName   		string    `json:"username"`
jwt.StandardClaims


}
func InitJwt(cfg *utils.Config){
	JwtKey = []byte(cfg.BlogJwt)
}


//生成TOKEN
func SetToken(username string,password string)(string,int)  {
	nowtime:= time.Now().Add(10*time.Hour)
	Jwt := Jwtstruct{
		username,
		jwt.StandardClaims{
			ExpiresAt: nowtime.Unix(),
			Issuer:"crcblog",

		},
	}
	result:=jwt.NewWithClaims(jwt.SigningMethodHS256,Jwt)
	token,err:=result.SignedString(JwtKey)
	if err!=nil{
		return "",utils.ERROR
	}

	return token,utils.SUCCESS
}
//验证TOKEN
func PassToken(t string) (*Jwtstruct,int) {
	token,err:=jwt.ParseWithClaims(t,&Jwtstruct{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if err!=nil{
		panic(err)
	}
	if key,_:= token.Claims.(*Jwtstruct);token.Valid{
		return key,utils.SUCCESS
	} else{
	return nil,utils.ERROR
	}

}
//jwt中间件
func JwtToken()gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		token:=ctx.Request.Header.Get("Authorization")
		code :=utils.SUCCESS
		if token == ""{
			code = utils.TOKEN_ERROR
			ctx.JSON(http.StatusOK,gin.H{
				"code":code,
				"message":utils.ErrorMsg(code),
			})
			ctx.Abort()
			return
		}
		checkToken :=strings.SplitN(token," ",2)
		if len(checkToken)!=2&&checkToken[0]!="Bearer"{
			code =utils.TOKEN_ERROR
			ctx.JSON(http.StatusOK,gin.H{
				"code":code,
				"message":utils.ErrorMsg(code),
			})
			ctx.Abort()
			return
		}
		key,code2:=PassToken(checkToken[1])

		if code2 == utils.ERROR{
			code =utils.TOKEN_ERROR
			ctx.JSON(http.StatusOK,gin.H{
				"code":code,
				"message":utils.ErrorMsg(code),
			})
			ctx.Abort()
			return
		}
		if time.Now().Unix()>key.ExpiresAt {
			code = utils.TOKEN_ERROR
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": utils.ErrorMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Set("username",key.UserName)
		ctx.Next()
	}
}
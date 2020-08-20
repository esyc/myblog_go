package main

import (
	"crcblog/middleware"
	"crcblog/models"
	"crcblog/router"
	"crcblog/utils"
	"github.com/gin-gonic/gin"
)



func main()  {


	cfg,err:=utils.GetConfig("./config/config.json")//读取配置信息

	if err !=nil{
		panic(err.Error())
	}
	models.InitDatabase(cfg)
	if err!=nil{
		panic(err.Error())
		return
	}
	middleware.InitJwt(cfg)
	gin.SetMode(cfg.BlogMode)

	engine :=gin.Default()//加载默认引擎
	engine.Use(middleware.Cors())
	registerRoute(engine)
	engine.Run(cfg.BlogHost+":"+cfg.BlogPort)

}
/*
*****注册路由*****
*/
func registerRoute(engine *gin.Engine)  {
	router.Router(engine)
}
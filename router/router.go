package router

import (
	"crcblog/controller"
	"crcblog/middleware"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine)  {
	//用户模块路由组
	userRoute:=engine.Group("/user")
	userRoute.Use(middleware.JwtToken())
	{
		userRoute.DELETE("delete/:name",controller.DeleteUser)
	}
	cateRoute:=engine.Group("/cate")
	cateRoute.Use(middleware.JwtToken())
	{
		cateRoute.POST("add",controller.AddCategory)
		cateRoute.PUT("edit/:id",controller.EditCategory)
		cateRoute.DELETE("delete/:id",controller.DeleteCategory)
	}
	postRoute:=engine.Group("/post")
	postRoute.Use(middleware.JwtToken())
	{
		postRoute.POST("add",controller.AddPost)
		postRoute.PUT("edit/:id",controller.EditPost)
		postRoute.DELETE("delete/:id",controller.DeletePost)

	}
	FileRoute:=engine.Group("/file")
	FileRoute.Use(middleware.JwtToken())
	{
		FileRoute.POST("uploadFile",controller.UploadFile)
		FileRoute.DELETE("deleteFile/:filename",controller.DeleteFile)
	}
	FileRoute2:=engine.Group("/file")
	FileRoute2.GET("img/:filename",controller.GetFile)
	userRoute2:=engine.Group("/user")
	{
		userRoute2.POST("login",controller.UserLogin)
		userRoute2.POST("add",controller.AddUser)
	}

	cateRoute2:=engine.Group("/cate")
	cateRoute2.GET("get",controller.GetAllCategory)
	postRoute2:=engine.Group("/post")
	{
		postRoute2.GET("get/number",controller.GetPostNum)
		postRoute2.GET("get", controller.GetAllPost)//全部获取
		postRoute2.GET("get/num", controller.GetPostByNum)//分页获取
		postRoute2.GET("get/cate/:cid", controller.GetPostFromCate)//分类获取
		postRoute2.GET("get/one/:id", controller.GetOnePost)//单个获取
	}
	//postRoute:=engine.Group("/post")
	//{
		//postRoute.GET("find",)
	//}
	//cateRoute:=engine.Group("/cate")
	//{
		//cateRoute.GET()
	//}
	otherRoute:=engine.Group("/other")
	{
		otherRoute.GET("get/zan",controller.GetZan)
		otherRoute.GET("get/msg",controller.GetMessage)
		otherRoute.GET("get/peo",controller.GetPeo)
		otherRoute.POST("add/msg",controller.AddMessage)
		otherRoute.POST("add/zan",controller.AddZan)
		otherRoute.POST("add/peo",controller.AddPeo)
	}

}

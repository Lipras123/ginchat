package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//statics
	r.Static("/asset", "asset/")
	// r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	r.LoadHTMLGlob("view/**/*")
	//首页
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	//用户模块
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser) //***23.3.20**
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)

	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendMsgUser", service.SendUserMsg)
	return r
}

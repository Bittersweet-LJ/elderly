package router

import (
	"elderly/controller"
	_ "elderly/docs"
	"elderly/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin 设置为发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//os.Setenv("NAME_OF_ENV_VARIABLE", "11")
	//注册swagger api路由信息
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	v1 := r.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)
	//获取温湿度数据列表
	v1.GET("/dht22", controller.GetDhtListHandler)
	//获取烟感灯状态数据列表
	v1.GET("/smooth_led", controller.GetLedStateListHandler)

	return r
}

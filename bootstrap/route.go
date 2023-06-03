package bootstrap

import (
	"github.com/gin-gonic/gin"
	"largo/config"
)

func SetupRoute(r *gin.Engine) {

	// 注册全局中间件
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	//注册路由
	config.RegisterRoutes(r)
}

package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"love-recycling-go/global"
	"love-recycling-go/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	cmsGroup := router.Group("/cms")
	routes.SetCmsGroupRoutes(cmsGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	err := r.Run(":" + global.App.Config.App.Port)
	if err != nil {
		fmt.Println("启动失败")
	}
}

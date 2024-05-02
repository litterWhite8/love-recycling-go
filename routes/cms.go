package routes

import (
	"github.com/gin-gonic/gin"
	"love-recycling-go/handler/cms"
	"love-recycling-go/middleware"
	"love-recycling-go/utils"
)

// SetCmsGroupRoutes  cms接口路由配置/*
func SetCmsGroupRoutes(router *gin.RouterGroup) {

	router.POST("/login", cms.Login)
	router.POST("/logout", cms.LogOut)

	router.Use(middleware.JWTAuth(utils.AppGuardName))
	router.GET("/getUserInfo", cms.GetUserInfo)
}

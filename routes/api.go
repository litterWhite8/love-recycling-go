package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"love-recycling-go/models"
	"love-recycling-go/utils"
	"net/http"
)

// SetApiGroupRoutes 接口测试样例 /*
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/ping", func(c *gin.Context) {
		var user models.SysUser
		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println(err)
			fmt.Println(user)
			c.JSON(http.StatusOK, gin.H{
				"error": utils.GetErrorMsg(user, err),
			})
			return
		}

		utils.SuccessNoData(c, "你是个好人")
	})
}

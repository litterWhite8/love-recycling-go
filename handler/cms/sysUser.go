package cms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"love-recycling-go/global"
	"love-recycling-go/models"
	"love-recycling-go/utils"
	"net/http"
)

func Login(c *gin.Context) {
	var loginUser models.SysUser
	if err := c.ShouldBind(&loginUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": utils.GetErrorMsg(loginUser, err),
		})
		return
	}

	var user models.SysUser
	err := global.App.DB.Where("user_name = ?", loginUser.UserName).First(&user).Error
	fmt.Println(loginUser)
	fmt.Println(user)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "用户不存在",
		})
		return
	}

	if user.Password != loginUser.Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "账号或密码错误",
		})
		return
	}
	tokenData, err, _ := utils.JwtService.CreateToken(utils.AppGuardName, user)

	utils.Success(c, tokenData, "登录成功")
}

func LogOut(c *gin.Context) {
	err := utils.JwtService.JoinBlackList(c.Request.Header.Get("Authorization"))
	if err != nil {
		utils.Fail(c, 500, "登出失败")
		return
	}
	utils.Success(c, nil, "登出成功")
}

func GetUserInfo(c *gin.Context) {
	var user models.SysUser
	err := global.App.DB.Where("user_name = ?", c.Query("userName")).First(&user).Error
	if err != nil {
		utils.Fail(c, 500, "用户不存在")
		return
	}
	utils.Success(c, user, "ok")

}

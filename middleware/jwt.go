package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"love-recycling-go/global"
	"love-recycling-go/models"
	"love-recycling-go/utils"
	"strconv"
	"time"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")

		if tokenStr == "" {
			utils.TokenFail(c)
			c.Abort()
			return
		}
		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.App.Token.Secret), nil
		})
		if err != nil || utils.JwtService.IsInBlacklist(tokenStr) {
			utils.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*utils.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			utils.TokenFail(c)
			c.Abort()
			return
		}

		//token续签
		if claims.ExpiresAt-time.Now().Unix() < global.App.Config.App.Token.RefreshGracePeriod {
			var user models.SysUser
			_ = global.App.DB.Where("user_id = ?", claims.Id).First(&user).Error
			tokenData, _, _ := utils.JwtService.CreateToken(GuardName, user)
			c.Header("new-token", tokenData.AccessToken)
			c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
			_ = utils.JwtService.JoinBlackList(tokenStr)
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}

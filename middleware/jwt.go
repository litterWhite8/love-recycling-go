package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"love-recycling-go/global"
	"love-recycling-go/utils"
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

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}

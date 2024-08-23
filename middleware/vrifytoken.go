package middleware

/*
 * @Author: <NAME>
 * 登录模块中间件
 * @Date: 2019-08-30 15:47:16
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	Config "gitee.com/under-my-umbrella/cloud/config"
)

// 登录验证中间件
func VrifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 token
		tokenString := c.Request.Header.Get("token")
		if tokenString == "" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "请登录",
			})
			c.Abort()
			return
		}

		// 解析 token
		token, err := jwt.ParseWithClaims(tokenString, &Config.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// 返回密钥
			return []byte("secretAtXiongHaiyin"), nil
		})

		if err != nil {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "无效的 Token",
			})
			c.Abort()
			return
		}

		// 验证 Claims
		if claims, ok := token.Claims.(*Config.CustomClaims); ok && token.Valid {
			// 将用户信息添加到上下文
			c.Set("id", claims.ID)
			c.Set("username", claims.Username)
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "无效的 Token",
			})
			c.Abort()
		}
	}
}

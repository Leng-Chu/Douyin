package jwt

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var invalidResp = common.Response{
	StatusCode: 1,
	StatusMsg:  "Invalid token",
}

// AuthWithLogin 鉴权中间件，为登录用户鉴权并通过解析token设置user_id
func AuthWithLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//没有获取到token
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, invalidResp)
			c.Abort()
			return
		}

		// 解析token
		token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, invalidResp)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*MyClaims)
		if !ok || !token.Valid || time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, invalidResp)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}

// AuthWithoutLogin 鉴权中间件，部分功能不要求登录，通过登录用户的token解析出user_id，未登录用户的user_id设置为0
func AuthWithoutLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		const defaultID uint = 0
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//没有获取到token
		if tokenStr == "" {
			c.Set("user_id", defaultID)
		} else {
			// 解析token
			token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil {
				c.Set("user_id", defaultID)
			} else {
				claims, ok := token.Claims.(*MyClaims)
				if !ok || !token.Valid || time.Now().Unix() > claims.ExpiresAt {
					c.Set("user_id", defaultID)
				} else {
					c.Set("user_id", claims.UserId)
				}
			}
		}
		c.Next()
	}
}

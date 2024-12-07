package middleware

import (
	"ctw-interview/model"
	"ctw-interview/response"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailedUnauthorized(gin.H{}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.FailWithDetailedUnauthorized(gin.H{}, err.Error(), c)
			c.Abort()
			return
		}
		user, err := model.GetUserById(claims.UserId)
		if err != nil {
			response.FailWithDetailedForbidden(gin.H{}, "用户信息查询失败", c)
			c.Abort()
			return
		}
		if user.Id <= 0 {
			response.FailWithDetailedForbidden(gin.H{}, "鉴权失败", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

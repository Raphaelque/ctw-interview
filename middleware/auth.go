package middleware

import (
	"github.com/gin-gonic/gin"
)

func UserAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		//authHelper(c)
	}
}

//func authHelper(c *gin.Context) {
//	session := sessions.Default(c)
//	username := session.Get("username")
//	role := session.Get("role")
//	id := session.Get("id")
//	status := session.Get("status")
//	fmt.Printf("id: %v, username: %v, role: %v\n\n", id, username, role)
//	useAccessToken := false
//	if username == nil {
//		// Check access token
//		accessToken := c.Request.Header.Get("Authorization")
//		if accessToken == "" {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"success": false,
//				"message": "无权进行此操作，未登录且未提供 access token",
//			})
//			c.Abort()
//			return
//		}
//		user := model.ValidateAccessToken(accessToken)
//		if user != nil && user.Username != "" {
//			if !validUserInfo(user.Username, user.Role) {
//				c.JSON(http.StatusOK, gin.H{
//					"success": false,
//					"message": "无权进行此操作，用户信息无效",
//				})
//				c.Abort()
//				return
//			}
//			// Token is valid
//			username = user.Username
//			role = user.Role
//			id = user.Id
//			status = user.Status
//			useAccessToken = true
//		} else {
//			c.JSON(http.StatusOK, gin.H{
//				"success": false,
//				"message": "无权进行此操作，access token 无效",
//			})
//			c.Abort()
//			return
//		}
//	}
//	if !useAccessToken {
//		// get header New-Api-User
//		apiUserIdStr := c.Request.Header.Get("New-Api-User")
//		idInt, ok := id.(int)
//		//apiUserIdStr := id.(int)
//		//if !ok {
//		//	apiUserIdStr = "0"
//		//}
//		fmt.Printf("id: %v, username: %v, role: %v, apiUserId: %v\n\n", id, username, role, apiUserIdStr)
//		if apiUserIdStr == "" && !ok {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"success": false,
//				"message": "无权进行此操作，请刷新页面或清空缓存后重试",
//			})
//			c.Abort()
//			return
//		}
//
//		var apiUserId int
//		if ok {
//			apiUserId = idInt
//		} else {
//
//		}
//
//		if ok {
//			apiUserId = idInt
//		} else {
//			var er error
//			apiUserId, er = strconv.Atoi(apiUserIdStr)
//			if er != nil {
//				c.JSON(http.StatusUnauthorized, gin.H{
//					"success": false,
//					"message": "无权进行此操作，登录信息无效，请重新登录",
//				})
//				c.Abort()
//				return
//
//			}
//		}
//
//		if id != apiUserId {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"success": false,
//				"message": "无权进行此操作，与登录用户不匹配，请重新登录",
//			})
//			c.Abort()
//			return
//		}
//	}
//	c.Set("username", username)
//	c.Set("role", role)
//	c.Set("id", id)
//	c.Set("group", session.Get("group"))
//	c.Set("use_access_token", useAccessToken)
//	c.Next()
//}

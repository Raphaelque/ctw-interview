package controller

import (
	"net/http"

	"ctw-interview/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//if !common.PasswordLoginEnabled {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "管理员关闭了密码登录",
	//		"success": false,
	//	})
	//	return
	//}
	//var loginRequest LoginRequest
	//err := json.NewDecoder(c.Request.Body).Decode(&loginRequest)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "无效的参数",
	//		"success": false,
	//	})
	//	fmt.Println("无效参数1")
	//	return
	//}
	//username := loginRequest.Username
	//password := loginRequest.Password
	//if username == "" || password == "" {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "无效的参数",
	//		"success": false,
	//	})
	//	fmt.Println("无效参数2")
	//	return
	//}
	//user := model.User{
	//	Username: username,
	//	Password: password,
	//}
	//err = user.ValidateAndFill()
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": err.Error(),
	//		"success": false,
	//	})
	//	fmt.Println("无效参数3")
	//	return
	//}
	setupLogin(nil, c)
}

func setupLogin(user *model.User, c *gin.Context) {
	//session := sessions.Default(c)
	//session.Set("id", user.Id)
	//session.Set("username", user.Username)
	//session.Set("role", user.Role)
	//session.Set("status", user.Status)
	//// raphael update 添加group进入session
	//session.Set("group", user.Group)
	//fmt.Printf("id: %v, username: %v, role: %v, status: %v\n", user.Id, user.Username, user.Role, user.Status)
	//err := session.Save()
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "无法保存会话信息，请重试",
	//		"success": false,
	//	})
	//	return
	//}
	//cleanUser := model.User{
	//	Id:          user.Id,
	//	Username:    user.Username,
	//	Email:       user.Email,
	//	DisplayName: user.DisplayName,
	//	Role:        user.Role,
	//	Status:      user.Status,
	//	Group:       user.Group,
	//	Amount:      user.Amount,
	//}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data":    nil,
	})
}

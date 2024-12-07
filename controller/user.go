package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ctw-interview/common"
	"ctw-interview/middleware"
	"ctw-interview/model"
	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerResponse struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func Register(c *gin.Context) {
	var user registerRequest
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "参数错误",
		})
		return
	}
	// 检查户名是否已存在
	exist, err := model.CheckUserNameExist(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "DB操作失败",
		})
		common.SysError(fmt.Sprintf("CheckUserNameExist error: %v", err))
		return
	}
	if exist {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名已存在",
		})
		return
	}
	dbUser := model.User{
		Username: user.Username,
		Password: user.Password,
	}
	// 插入用户
	err = dbUser.CreateUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "注册成功",
	})
	return
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req loginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "无效的参数",
			"success": false,
		})
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}
	dbUser, err := user.CheckUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	setupLogin(dbUser, c)
}

type loginResponse struct {
	Token  string `json:"token"`
	UserID int64  `json:"userId"`
}

func setupLogin(user *model.User, c *gin.Context) {
	// 生成token
	jwtToken, err := middleware.GenerateToken(user.Id)
	// 更新token
	user.Token = jwtToken
	user, err = user.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data": loginResponse{
			Token:  user.Token,
			UserID: user.Id},
	})
}

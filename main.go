package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ctw-interview/common"
	"ctw-interview/middleware"
	"ctw-interview/model"
	"ctw-interview/router"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// 读取 .env 文件
	err := godotenv.Load(".env")
	if err != nil {
		common.SysLog("Can't load .env file")
	}

	common.SetupLogger()
	common.SysLog("CTW interview " + " started")
	//if os.Getenv("GIN_MODE") != "debug" {
	//	gin.SetMode(gin.DebugMode)
	//}
	if common.DebugEnabled {
		common.SysLog("running in debug mode")
		gin.SetMode(gin.DebugMode)
	}

	// Initialize SQL Database
	err = model.InitDB()
	if err != nil {
		common.FatalLog("failed to initialize database: " + err.Error())
	}

	// Initialize HTTP server
	server := gin.New()
	server.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		common.SysError(fmt.Sprintf("panic detected: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"message": fmt.Sprintf("Panic detected, error: %v.", err),
				"type":    "CTW_panic",
			},
		})
	}))

	server.Use(middleware.RequestId())
	middleware.SetUpLogger(server)

	router.SetRouter(server)
	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(*common.Port)
	}
	err = server.Run(":" + port)
	if err != nil {
		common.FatalLog("failed to start HTTP server: " + err.Error())
	}
}

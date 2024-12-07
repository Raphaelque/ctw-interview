package common

import (
	"os"
	"time"
)

var DebugEnabled = os.Getenv("DEBUG") == "true"

const (
	RequestIdKey = "X-CTW-Request-Id"
)

// 速率限制配置
var (
	GlobalApiRateLimitEnable   = GetEnvOrDefaultBool("GLOBAL_API_RATE_LIMIT_ENABLE", true)
	GlobalApiRateLimitNum      = GetEnvOrDefault("GLOBAL_API_RATE_LIMIT", 180)
	GlobalApiRateLimitDuration = int64(GetEnvOrDefault("GLOBAL_API_RATE_LIMIT_DURATION", 180))

	UploadRateLimitNum            = 10
	UploadRateLimitDuration int64 = 60

	DownloadRateLimitNum            = 10
	DownloadRateLimitDuration int64 = 60
)

var RateLimitKeyExpirationDuration = 20 * time.Minute

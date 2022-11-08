package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"[%s] [%s] %s %s %d %s %s\n",
			params.TimeStamp.Format(time.RFC3339),
			params.ClientIP,
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
		)
	})
}

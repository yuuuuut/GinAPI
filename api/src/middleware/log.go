package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Loging(c *gin.Context) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	userAgent := c.GetHeader("User-Agent")
	c.Next()

	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("ua", userAgent),
		zap.Int("status", c.Writer.Status()),
		zap.String("time", time.Now().Format("2006/01/02 15:04:05")),
	)
}

package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/willsmil/lover/pkg/utils/logger"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		trace := uuid.New().String()
		c.Set("trace", trace)
		logger.Logger.Info("trace: %s, request: %s", trace, c.Request.RequestURI)
		c.Next()
		logger.Logger.Info("trace %s, cost: %s", trace, time.Since(start).String())
	}
}

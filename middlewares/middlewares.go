package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()
		duration := time.Since(startTime)

		logger.WithFields(logrus.Fields{
			"Method":   c.Request.Method,
			"url":      c.Request.URL,
			"duration": duration,
		}).Info("Request Recived")
	}
}

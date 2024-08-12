package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var token = "some-token"

func Middleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		logger.WithFields(logrus.Fields{
			"Method": c.Request.Method,
			"url":    c.Request.URL,
		}).Info("Request Recived")
	}
}

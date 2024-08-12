package main

import (
	"Demo/handlers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MyHandler struct {
	logger *logrus.Logger
}

func NewHandlerWithLogger(logger *logrus.Logger) *MyHandler {
	return &MyHandler{logger: logger}
}

func main() {
	logger := logrus.New()
	router := gin.New()
	router.Use(handlers.Middleware(logger))
	myHandler := NewHandlerWithLogger(logger)
	router.GET("/hello", myHandler.Hello)
	router.Run(":8080")
}

func (h MyHandler) Hello(c *gin.Context) {
	h.logger.Info("dependency injected")
	c.JSON(200, "Hello")
}

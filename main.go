package main

import (
	handlers "Demo/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// struct to add dependency that our application holds
type MyHandler struct {
	logger *logrus.Logger
	// here new dependency can be added like tracing or other
}

// cretes new object and initialize
func NewHandlerWithLogger(logger *logrus.Logger) *MyHandler {
	return &MyHandler{logger: logger}
}

func main() {
	logger := logrus.New()
	router := gin.New()
	// call the middleware
	router.Use(handlers.LogingMiddleware(logger))
	myHandler := NewHandlerWithLogger(logger)
	router.GET("/hello", myHandler.Hello)
	router.Run(":8080")
}

func (h MyHandler) Hello(c *gin.Context) {
	h.logger.Info("dependency injected")
	c.JSON(200, "Hello")
}

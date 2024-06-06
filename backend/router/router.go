package router

import (
	"backend/internal/rabbitmq"
	"backend/router/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(hosts map[string]struct{}) *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Use(middleware.Security(hosts))

	conn, ch, err := rabbitmq.SetupRabbitMQ()
	if err != nil {
		panic(err)
	}

	return r
}

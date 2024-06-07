package router

import (
	"backend/db"
	"backend/internal/weather"
	"backend/router/middleware"
	"github.com/gin-gonic/gin"
	throttle "github.com/s12i/gin-throttle"
)

func WeatherRouter(hosts map[string]struct{}) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	conn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	maxEventPerSec := 1000
	maxBurstSize := 20

	r.Use(throttle.Throttle(maxEventPerSec, maxBurstSize))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Weather Service Pong!"})
	})
	repo := weather.NewRepository(conn.GetClient())
	service := weather.NewService(repo)
	handler := weather.NewHandler(service)

	r.Use(middleware.Security(hosts))

	protectedGroup := r.Group("/protected")
	jwtMiddleware := middleware.NewJWTMiddleware()
	{
		protectedGroup.Use(jwtMiddleware.Authorize())
		protectedGroup.GET("/", handler.FindWeather)
		protectedGroup.GET("/", handler.UpdateWeather)
	}

	return r
}

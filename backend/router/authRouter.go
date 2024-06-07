package router

import (
	"backend/db"
	"backend/internal/auth"
	"backend/router/middleware"
	"github.com/gin-gonic/gin"
	throttle "github.com/s12i/gin-throttle"
)

func AuthRouter(hosts map[string]struct{}) *gin.Engine {
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
		c.JSON(200, gin.H{"message": "Auth Service Pong!"})
	})
	repo := auth.NewRepository(conn.GetClient())
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)

	r.Use(middleware.Security(hosts))
	publicGroup := r.Group("/")
	{
		publicGroup.POST("/login", handler.Login)
		publicGroup.POST("/register", handler.Register)
	}
	protectedGroup := r.Group("/protected")
	jwtMiddleware := middleware.NewJWTMiddleware()
	{
		protectedGroup.Use(jwtMiddleware.Authorize())
		protectedGroup.GET("/", handler.GetUser)
		protectedGroup.GET("/", handler.Logout)
	}

	return r
}

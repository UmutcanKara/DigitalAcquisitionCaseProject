package router

import (
	"backend/db"
	"backend/internal/auth"
	"backend/router/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	throttle "github.com/s12i/gin-throttle"
	"time"
)

func AuthRouter(hosts map[string]struct{}) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	conn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	repo := auth.NewRepository(conn.GetClient())
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)

	maxEventPerSec := 1000
	maxBurstSize := 20

	// Configure CORS middleware
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://172.17.0.2:5173"}, // Replace with your allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply the CORS middleware to the router
	r.Use(cors.New(corsConfig))

	r.Use(throttle.Throttle(maxEventPerSec, maxBurstSize))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Auth Service Pong!"})
	})

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
		protectedGroup.POST("/logout", handler.Logout)
	}

	return r
}

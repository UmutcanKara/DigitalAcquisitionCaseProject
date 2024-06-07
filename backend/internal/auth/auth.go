package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Auth struct {
	Username string
	Password string
	Hometown string
}

type Repository interface {
	login(username, password string) error
	register(username, password, hometown string) error
	getUser(username string) (bson.M, error)
}

type Service interface {
	login(ctx context.Context, req LoginReq) error
	register(ctx context.Context, req RegisterReq) error
	getUser(ctx context.Context, username string) (bson.M, error)
	createJWTToken(ctx context.Context, username string) (string, error)
}

type Handler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetUser(c *gin.Context)
	Logout(c *gin.Context)
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hometown string `json:"hometown"`
}
type RegisterRes struct {
	Token string `json:"token"`
}
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginRes struct {
	Token string `json:"token"`
}

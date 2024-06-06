package auth

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Auth struct {
	Username string
	Password string
	Hometown string
}

type Repository interface {
	login(username, password string) (bson.M, error)
	register(username, password, hometown string) error
}

type Service interface {
	login(ctx context.Context, req LoginReq) (LoginRes, error)
	register(ctx context.Context, req RegisterReq) (RegisterRes, error)
}

type Handler interface {
	Login(req LoginReq) string
	Register(req RegisterReq) string
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

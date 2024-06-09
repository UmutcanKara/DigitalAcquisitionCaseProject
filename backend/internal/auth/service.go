package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(r Repository) Service {
	return &service{r, time.Duration(5) * time.Second}
}

func (s *service) login(ctx context.Context, req LoginReq) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.Repository.login(req.Username, req.Password); err != nil {
		return err
	}

	return nil
}

func (s *service) register(ctx context.Context, req RegisterReq) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.Repository.register(req.Username, req.Password, req.Hometown); err != nil {
		return err
	}

	return nil
}

func (s *service) getUser(ctx context.Context, username string) (bson.M, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	result, err := s.Repository.getUser(username)
	if err != nil {
		return result, err
	}
	result["password"] = ""
	return result, nil
}

func (s *service) createJWTToken(ctx context.Context, username string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    username,
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(24) * time.Hour)),
	})
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

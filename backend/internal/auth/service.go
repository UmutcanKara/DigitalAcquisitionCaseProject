package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
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

func (s *service) login(ctx context.Context, req LoginReq) (LoginRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if _, err := s.Repository.login(req.Username, req.Password); err != nil {
		return LoginRes{}, err
	}

	ss, err := createJWTToken(req.Username)
	if err != nil {
		return LoginRes{}, err
	}

	return LoginRes{ss}, nil
}

func (s *service) register(ctx context.Context, req RegisterReq) (RegisterRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.Repository.register(req.Username, req.Password, req.Hometown); err != nil {
		return RegisterRes{}, err
	}

	ss, err := createJWTToken(req.Username)
	if err != nil {
		return RegisterRes{}, err
	}
	return RegisterRes{ss}, nil
}

func createJWTToken(username string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    username,
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

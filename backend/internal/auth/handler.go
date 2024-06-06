package auth

import (
	"context"
	"github.com/goccy/go-json"
)

type handler struct {
	Service
}

func NewHandler(s Service) Handler { return &handler{s} }

func (h *handler) Login(req LoginReq) string {
	ctx := context.Background()

	res, err := h.Service.login(ctx, req)
	if err != nil {
		return err.Error()
	}
	marshalBytes, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(marshalBytes)
}

func (h *handler) Register(req RegisterReq) string {
	ctx := context.Background()

	res, err := h.Service.register(ctx, req)
	if err != nil {
		return err.Error()
	}
	marshalBytes, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(marshalBytes)
}

package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	Service
}

func NewHandler(s Service) Handler { return &handler{s} }

func (h *handler) Login(c *gin.Context) {
	req := LoginReq{}
	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Service.getUser(ctx, req.Username)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Invalid credentials"})
		return
	}

	err = h.Service.login(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ss, err := h.Service.createJWTToken(ctx, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("Authorization", ss, 60*15, "/", "", false, true)
	c.JSON(http.StatusOK, LoginRes{ss})
}

func (h *handler) Register(c *gin.Context) {
	req := RegisterReq{}
	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Service.getUser(ctx, req.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	err = h.Service.register(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ss, err := h.Service.createJWTToken(ctx, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("Authorization", ss, 60*15, "/", "", false, true)
	c.JSON(http.StatusOK, RegisterRes{ss})
}

func (h *handler) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
}

func (h *handler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	username := c.Query("username")

	result, err := h.Service.getUser(ctx, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, result)
}

package handlers

import (
	"net/http"

	"github.com/20ritiksingh/hospital-app/internal/mapper"
	"github.com/20ritiksingh/hospital-app/internal/openapi"
	"github.com/gin-gonic/gin"
)

func (h *APIHandler) Signup(c *gin.Context) {
	var req openapi.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user := mapper.MapSignupReqestToUser(&req)
	token, err := h.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, openapi.LoginResponse{
		AccessToken: &token,
	})
}

func (h *APIHandler) Login(c *gin.Context) {
	var req openapi.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user := mapper.MapLoginRequestToUser(&req)
	token, err := h.authService.Login(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, openapi.LoginResponse{
		AccessToken: &token,
	})
}

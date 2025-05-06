// controllers/auth_controller.go
package controllers

import (
	"net/http"
	"todo-case/utils"
	"todo-case/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{UserService: userService}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ctl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Veriler geçersiz"})
		return
	}

	user, ok := ctl.UserService.ValidateCredentials(req.Username, req.Password)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı adı veya şifre hatalı"})
		return
	}

	token, err := utils.GenerateKey(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token oluşturulamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

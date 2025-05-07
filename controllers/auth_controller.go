// controllers/auth_controller.go
package controllers

import (
	"net/http"
	"todo-case/services"
	"todo-case/utils"

	"todo-case/models"

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
type LoginResponse struct {
	Token string `json:"token" example:"token123"`
}



// Login godoc
// @Summary      Giriş yap
// @Description  Kullanıcı adı ve şifre ile giriş yapar
// @Tags         auth
// @Accept       json
// @Param        LoginRequest  body      LoginRequest  true  "Yeni görev listesi"
// @Produce      json
// @Success      200  {array}  models.ToDo
// @Failure      500  {object}  models.ErrorResponse
// @Router       /login [post]
func (ctl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Veriler geçersiz"})
		return
	}

	user, ok := ctl.UserService.ValidateCredentials(req.Username, req.Password)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Geçersiz kullanici adi veya şifre"})
		return
	}

	token, err := utils.GenerateKey(user.ID, user.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Token oluşturulamadi"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

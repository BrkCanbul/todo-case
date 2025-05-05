package controllers

import (
	"net/http"
	"todo-case/utils"

	"github.com/gin-gonic/gin"
)

type login struct {
	UserId uint "json:user_id"
}

func Login(c *gin.Context) {
	var input login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateKey(input.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't create token", "details": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

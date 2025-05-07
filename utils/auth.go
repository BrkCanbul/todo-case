package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("secretKey")

type Claims struct {
	UserId   uint
	UserType string
	jwt.RegisteredClaims
}

func GenerateKey(userId uint,role string) (string, error) {
	expTime := time.Now().Add(24 * time.Hour)
	if role != "admin" && role != "user" {
		// make sure if invalid role is given the program will error
		return "", jwt.ErrTokenInvalidClaims

	} 
	claims := &Claims{
		UserId: userId,
		UserType: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func Validate(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) { return key, nil })
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Yetkisiz"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := Validate(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz token"})
			return
		}

		c.Set("user_id", claims.UserId)
		fmt.Printf("User ID: %d\n", claims.UserId)
		fmt.Printf("User Type: %s\n", claims.UserType)
		c.Set("user_type", claims.UserType)

		c.Next()
	}
}

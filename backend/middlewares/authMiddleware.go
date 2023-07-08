package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	secretKey := os.Getenv("SECRETKEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// Verify the token
	if err != nil || !token.Valid {
		fmt.Println("Token In valid")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Set the user information from the token in the Gin context
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"]
	AdminLoginEmail := os.Getenv("AdminLoginEmail")
	if AdminLoginEmail == email {
		c.Set("admin", 1)
		c.Set("email", email)
		c.Next()
	} else {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	fmt.Println("Authorized")
	c.Set("email", user.Email)
	c.Set("name", user.Name)
	c.Set("id", user.ID)
	c.Next()
    }
}

package middlewares

import (
	"net/http"
	"os"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func isAuth(c *gin.Context) {
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Set the user information from the token in the Gin context
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"]
	var user models.User 
	result := database.DB.Where("email = ?",email).First(&user)
    if(result != nil ){
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("email",user.Name)
	c.Set("name",user.Name)
	c.Set("id",user.ID)
	c.Next()
}

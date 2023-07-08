package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context){
     var user models.SignUps
	 err := c.ShouldBindJSON(&user)
     if err != nil {
		c.JSON(400,gin.H{
			"error":"Looks like Credentials are Wrong!!",
		})
	 } else {
		var temp models.User
		   result := database.DB.Where("email = ?",user.Email).First(&temp)
		   if(result.Error == nil){
			c.JSON(400,gin.H{
				"error":"User with this email already exists",
			})
			return 
		   }
		   var temp_user models.SignUps
		   res := database.DB.Where("email = ?",user.Email).First(&temp_user)
		   if res.Error == nil {
				user.ID = temp_user.ID
		   }
           currentTime := time.Now()
		   timeString := currentTime.Format(time.RFC3339)
		   timeString=timeString+user.Email+user.Password
		   hash := sha256.Sum256([]byte(timeString))
		   token := hex.EncodeToString(hash[:])
		   url :=  "http://localhost:8080/signup/verify/" + token 
		   user.VerifyToken = token
		   AdminEmail := os.Getenv("AdminEmail")
		   AdminPassword := os.Getenv("AdminPassword")
		   email := &models.Email{
               Subject: "Verfication Email",
			   Body : "Click the Link for verification of your account \r\n"+url+"\r\n",
			   SenderEmail: AdminEmail,
			   Password: AdminPassword,
			   ReceiverEmail: user.Email,
		   }
		   database.DB.Save(&user)
		   err := SendEmail(email)
		   if err != nil {
			c.JSON(400,gin.H{
				"error" : "Error in Registering the account\r\n Please Try Again!",
			})
		   } else {
			c.JSON(200, gin.H{
				"message":"Email Sent! Verify your account!",
			})
		   }
	 }

}

func VerifySignUps(c *gin.Context){
	token := c.Param("token")
	var signup models.SignUps
	result := database.DB.Where("verify_token = ?",token).First(&signup)
    if result.Error != nil {
		c.JSON(400,gin.H{
			"error":"Invalid Token!!",
		})
	} else {
		var user models.User
		user.Email = signup.Email
		user.Password = signup.Password
		user.ProfileImage_URL = signup.ProfileImage_URL
		user.Name = signup.Name
		result := database.DB.Create(&user)
		if result.Error != nil {
			c.JSON(500,gin.H{
				"error":"Unable to Create User! Try Again!!",
			})
		} else {
            database.DB.Delete(signup)
			c.JSON(200,gin.H{
				"message":"Successfully Created User! Try SignIn",
			})
		}
	}

}

func LogIn(c *gin.Context){
	type login struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	var x login
	c.ShouldBindJSON(&x)
	secretKey := os.Getenv("SECRETKEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] =  x.Email
    // add expiration later

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Login"})
		return
	}
	var user models.User
	result := database.DB.Where("email = ?",x.Email).First(&user)
    if result.Error != nil {
		c.JSON(400,gin.H{
			"error":"Invalid Email!",
		})
		return 
	}
	if x.Password == user.Password {
	c.JSON(200,gin.H{
		"token": tokenString,
		"message":"LogIn Successful",
	})
	return 
    }
	c.JSON(400,gin.H{
		"error":"Invalid Password or Email",
	})
	 
}
package routes

import (
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutesForProducts(router *gin.Engine){
// 	router.GET("/products",GetProducts);
// 	router.GET("/products/:id",GetProductDetails)
    router.GET("/getReviews/:id",handlers.GetReviews)
}

func SetUpRoutesForUsers(router *gin.Engine){
	// router.GET("/cart",GetCart);
	router.GET("/buy",handlers.BuyProducts);
	router.POST("/addReview/:id",handlers.AddReview)
	router.DELETE("/deleteReview/:id",handlers.RemoveReview)
}

func SetUpRoutesForAuth(router *gin.Engine){
	router.POST("/signin",handlers.LogIn);
	router.POST("/signup",handlers.SignUp);
	router.GET("/signup/verify/:token",handlers.VerifySignUps)
	//router.GET("/forgotpassword",HandleForgotPassword);
}

func SetUpRoutesForAdmin(router *gin.Engine){
	router.POST("/addProduct",handlers.AddProduct);
	router.DELETE("/removeProduct",handlers.RemoveProduct);
}

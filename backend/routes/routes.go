package routes

import (
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/handlers"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUpRoutesForProducts(router *gin.Engine){
 	router.GET("/products",handlers.GetProducts);
 	router.GET("/products/:id",handlers.GetProductDetails)
    router.GET("/getReviews/:id",handlers.GetReviews)
}

func SetUpRoutesForUsers(router *gin.Engine){
	// router.GET("/cart",GetCart);
	router.GET("/buy",middlewares.IsAuth,handlers.BuyProducts);
	router.POST("/addReview/:id",middlewares.IsAuth,handlers.AddReview)
	router.DELETE("/deleteReview/:id",middlewares.IsAuth,handlers.RemoveReview)
}

func SetUpRoutesForAuth(router *gin.Engine){
	router.POST("/signin",handlers.LogIn);
	router.POST("/signup",handlers.SignUp);
	router.GET("/signup/verify/:token",handlers.VerifySignUps)
	//router.GET("/forgotpassword",HandleForgotPassword);
}

func SetUpRoutesForAdmin(router *gin.Engine){
	router.POST("/addProduct",middlewares.IsAuth,handlers.AddProduct);
	router.DELETE("/removeProduct",middlewares.IsAuth,handlers.RemoveProduct);
}

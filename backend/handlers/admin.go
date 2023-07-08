package handlers

import (

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(200,gin.H{
			"error":"Wrong JSON Format",
		}) 
	}
	product.Total_reviews=0;
	product.TotalRating=0;
	result := database.DB.Create(&product)
	if(result.Error != nil){
	c.JSON(200,gin.H{
			"error":"Unsuccessful to create a product",
		})
	} else {
		c.JSON(200,gin.H{
        "message":"Successful Creation of Product",
		})
	}	
}

func RemoveProduct(c *gin.Context){
      idStr := c.Param("id")
	  result := database.DB.Where("id = ?",idStr).Delete(&models.Product{})
	  if(result.Error != nil){
		c.JSON(400,gin.H{
			"error":"Error on Deleting Product",
		})  
	  } else {
          c.JSON(200,gin.H{
			"message":"Deleted Product Successfully!",
		  })
	  }
}
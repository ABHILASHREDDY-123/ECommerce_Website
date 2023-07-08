package handlers

import (
	"strconv"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/gin-gonic/gin"
)
func AddReview(c *gin.Context){
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
    var product models.Product
    result := database.DB.Where("id = ?",idStr).First(&product)
	if(err!=nil || result.Error != nil){
		c.JSON(400,gin.H{
			"error":"Can't add a Review to this product",
		})
	} else {
		var review_obj models.Review 
		err := c.ShouldBindJSON(&review_obj)
		if err != nil {
			c.JSON(400,gin.H{
				"error":"Can't add a Review to this product",
			}) 
		} else {
			review_obj.ProductID = uint(id)
			result := database.DB.Create(&review_obj)
			if result.Error != nil {
				c.JSON(400,gin.H{
					"error":"Can't add a Review to this product",
				}) 
			} else {
				product.Total_reviews++
				product.TotalRating+=int64(review_obj.Rating)
				database.DB.Save(product)
				c.JSON(200,gin.H{
					"message":"Added Review Successfully",
				})
			}
		}
	}
}

func RemoveReview(c *gin.Context){
   idStr := c.Param("id")
   id, err := strconv.Atoi(idStr)
   if(err != nil){
	c.JSON(200,gin.H{
		"error":"Review Id Error",
	})
   } else {
	var review_obj models.Review
	 result := database.DB.Where("id = ?",id).First(&review_obj)
	 if(result.Error != nil){
		c.JSON(200,gin.H{
			"error":"Already Deleted!",
		})
	 } else {
		var product models.Product
		result := database.DB.Where("id = ?",review_obj.ProductID).First(&product)
		if(result.Error != nil){
			c.JSON(200,gin.H{
				"error":"Error in Deleting!",
			})
		} else {
			product.TotalRating = product.TotalRating - int64(review_obj.Rating)
            product.Total_reviews--
            database.DB.Save(product)
			database.DB.Where("id = ?",review_obj.ID).Delete(&models.Review{})
		c.JSON(200,gin.H{
			"message":"Deleted Review Successfully",
		})
	}
	 }
   }
}
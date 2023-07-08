package handlers

import (
	"strconv"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Product Id",
		})
	}
	var review_objs []models.Review
	result := database.DB.Where("product_id=?", id).Find(&review_objs)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Cannot get reviews at this moment!",
		})
	} else {
		c.JSON(200, gin.H{
			"content": review_objs,
			"message": "Successfully Retreived Reviews!",
		})
	}

}

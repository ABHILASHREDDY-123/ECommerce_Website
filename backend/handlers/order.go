package handlers

import (
	"encoding/json"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/gin-gonic/gin"
)

func BuyProducts(c *gin.Context) {
	var order models.Order
	id := c.GetUint("id")
	type address struct{
		address string `json:"address"`
	}
	var x address
	c.ShouldBindJSON(&x)
	var user models.User
	res := database.DB.Where("id = ?",id).First(&user)
	if(res.Error!=nil){
       c.JSON(500,gin.H{
		"error":"Currently Unable to Place Order",
	   })
	} else {
		var Cart models.Cart
		var OrderObjs []models.Obj
		json.Unmarshal(user.Cart,&Cart)
		var obj models.Obj
		var prod models.Product
		for _, val := range Cart.Objs {
			obj.Quantity = uint64(val.Quantity)
			obj.ProductID = val.ProductId

			res := database.DB.Where("id = ?",obj.ProductID).First(&prod)
			if res.Error != nil {
				c.JSON(400,gin.H{
					"error":"Invalid Cart Items!!",
				})
				c.Abort()
				return 
			}
			obj.CostPerUnit = uint64(prod.Price)
			OrderObjs = append(OrderObjs, obj)
		}
		bytes,err := json.Marshal(OrderObjs)
		if err != nil {
			c.JSON(400,gin.H{
				"error":"Invalid Cart Items!!",
			})
			c.Abort()
			return
		}
		order.Items  = bytes 
		order.UserID = id
		order.Address  = x.address
		result := database.DB.Create(&order)
		if result.Error != nil {
			c.JSON(400,gin.H{
				"error":"Unable to Place Order",
			})
			c.Abort()
			return
		} else {
			c.JSON(200,gin.H{
             "message":"Successfully Placed Order",
			})
		}
	}
}


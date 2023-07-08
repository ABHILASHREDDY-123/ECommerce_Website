package main

import (
	"net/http"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/database"
	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()

	router.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK, "WELCOME HOME!")
	})
	routes.SetUpRoutesForAdmin(router)
	routes.SetUpRoutesForUsers(router)
	routes.SetUpRoutesForProducts(router)
	routes.SetUpRoutesForAuth(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

package main

import (
	auth "main/middlewares/auth"
	routes "main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	private := router.Group("/")
	private.Use(auth.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run(":8083")
}

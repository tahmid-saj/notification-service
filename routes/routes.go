package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// email
	server.POST("/email/verify")
	server.POST("/email/send")

	// notification
	server.GET("/topics")
	server.POST("/topics")
	server.GET("/topics/")
}
package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// email
	// VerifyEmail
	server.POST("/email/verify", verifyEmail)
	// SendEmail
	server.POST("/email/send", sendEmail)

	// notification
	// ListTopics
	server.GET("/topics")
	// CreateTopic
	server.POST("/topics")
	// ListSubscriptions
	server.GET("/topics/:topicARN")
	// SubscribeToTopic
	server.POST("/topics/:topicARN")
	// PublishMessageToAllTopicSubscribers
	server.POST("/topics/:topicARN/publish")
}
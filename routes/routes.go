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
	server.GET("/topics", getTopics)
	// CreateTopic
	server.POST("/topics/:topicName", postCreateTopic)
	// ListSubscriptions
	server.GET("/topics/:topicARN", getTopicSubscriptions)
	// SubscribeToTopic
	server.POST("/topics/:topicARN", postSubscribeToTopic)
	// PublishMessageToAllTopicSubscribers
	server.POST("/topics/:topicARN/publish", postPublishMessageToAllTopicSubscribers)
}
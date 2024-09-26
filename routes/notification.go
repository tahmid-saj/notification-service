package routes

import (
	"net/http"
	"notification-service/models"

	"github.com/gin-gonic/gin"
)

// notification
func getTopics(context *gin.Context) {
	resTopics, err := models.ListTopics()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch topics"})
		return
	}

	context.JSON(http.StatusOK, resTopics)
}

func postCreateTopic(context *gin.Context) {
	topicName := context.Param("topicName")

	resTopicCreated, err := models.CreateTopic(topicName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create topic"})
		return
	}

	context.JSON(http.StatusOK, resTopicCreated)
}

func getTopicSubscriptions(context *gin.Context) {
	topicARN := context.Param("topicARN")

	resTopicSubscriptions, err := models.ListSubscriptions(topicARN)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch subscriptions"})
		return
	}

	context.JSON(http.StatusOK, resTopicSubscriptions)
}

func postSubscribeToTopic(context *gin.Context) {
	topicARN := context.Param("topicARN")
	var subscriptionInput models.SubscriptionInput

	err := context.ShouldBindJSON(subscriptionInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	subscription := models.Subscription{
		TopicARN: topicARN,
		Email: subscriptionInput.Email,
	}

	resSubscription, err := models.SubscribeToTopic(subscription)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not subscribe to topic"})
		return
	}

	context.JSON(http.StatusOK, resSubscription)
}

func postPublishMessageToAllTopicSubscribers(context *gin.Context) {
	topicARN := context.Param("topicARN")
	var messageInput models.MessageInput

	err := context.ShouldBindJSON(messageInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	publishMessage := models.PublishMessage{
		TopicARN: topicARN,
		Message: messageInput.Message,
	}

	resPublishMessage, err := models.PublishMessageToAllTopicSubscribers(publishMessage)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not publish messages"})
		return
	}

	context.JSON(http.StatusOK, resPublishMessage)
}
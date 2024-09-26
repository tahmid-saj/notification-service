package main

import (
	// "notification-service/routes"
	// "os"

	// "github.com/gin-gonic/gin"
	"fmt"
	"notification-service/notification"

	"github.com/joho/godotenv"
)

func main() {
	// VerifyEmail
	// err := email.VerifyEmail("", "")
	// if err != nil {
	// 	return
	// }

	// err = email.SendEmail("", "", "subject", "text1", "test2", "UTF-8")
	// if err != nil {
	// 	return
	// }

	// ListTopics
	// topics, err := notification.ListTopics()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(topics)

	// ListSubscriptions
	// topicARN := ""
	// subscriptions, err := notification.ListSubscriptions(&topicARN)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(subscriptions)

	// SubscribeToTopic
	// email := ""
	// topicARN := ""
	// subscription, err := notification.SubscribeToTopic(&email, &topicARN)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(subscription)

	// message := ""
	// topicARN := ""
	// publishOutput, err := notification.PublishMessageToAllTopicSubscribers(&message, &topicARN)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(publishOutput)

	godotenv.Load()

	// server := gin.Default()

	// routes.RegisterRoutes(server)

	// server.Run(os.Getenv("PORT"))
}
package models

import (
	"notification-service/notification"
)

type Subscription struct {
	Email    string `json:"email"`
	TopicARN string `json:"topicARN"`
}

type PublishMessage struct {
	Message  string `json:"message"`
	TopicARN string `json:"topicARN"`
}

type SubscriptionInput struct {
	Email    string `json:"email"`
}

type MessageInput struct {
	Message string `json:"message"`
}

type TopicInput struct {
	TopicName string `json:"topicName"`
}

// notification
func ListTopics() (*Response, error) {
	topics, err := notification.ListTopics()
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: topics,
	}, nil
}

func CreateTopic(topicName string) (*Response, error) {
	createdTopic, err := notification.CreateTopic(topicName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: createdTopic,
	}, nil
}

func ListSubscriptions(topicARN string) (*Response, error) {
	subscriptions, err := notification.ListSubscriptions(&topicARN)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: subscriptions,
	}, nil
}

func SubscribeToTopic(subscriptionInput Subscription) (*Response, error) {
	subscription, err := notification.SubscribeToTopic(&subscriptionInput.Email, &subscriptionInput.TopicARN)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: subscription,
	}, nil
}

func PublishMessageToAllTopicSubscribers(publishMessageInput PublishMessage) (*Response, error) {
	publishedMessage, err := notification.PublishMessageToAllTopicSubscribers(&publishMessageInput.Message, &publishMessageInput.TopicARN)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: publishedMessage,
	}, nil
}
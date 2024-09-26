package notification

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func ListTopics() ([]*sns.Topic, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var topics []*sns.Topic
	topics = append(topics, result.Topics...)

	return topics, nil
}

func CreateTopic(topicName string) (*sns.CreateTopicOutput, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(topicName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func ListSubscriptions(topicPtr *string) ([]*sns.Subscription, error) {
	if *topicPtr == "" {
		fmt.Println("You must supply a topic ARN")
		return nil, errors.New("must supply email and topic")
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)
	var previousToken *string

	result, err := svc.ListSubscriptionsByTopic(&sns.ListSubscriptionsByTopicInput{
		NextToken: previousToken,
		TopicArn: topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result.Subscriptions, nil
}

func SubscribeToTopic(emailPtr *string, topicPtr *string) (*sns.SubscribeOutput, error) {
	if *emailPtr == "" || *topicPtr == "" {
		fmt.Println("You must supply an email address and topic ARN")
		return nil, errors.New("must supply email and topic")
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              emailPtr,
		Protocol:              aws.String("email"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              topicPtr,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func PublishMessageToAllTopicSubscribers(messagePtr *string, topicPtr *string) (*sns.PublishOutput, error) {

	if *messagePtr == "" || *topicPtr == "" {
		fmt.Println("You must supply a message and topic ARN")
		return nil, errors.New("must supply message and topic")
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:  messagePtr,
		TopicArn: topicPtr,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}
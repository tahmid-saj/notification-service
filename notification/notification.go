package notification

import (
	"errors"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func ListTopics() error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, t := range result.Topics {
		fmt.Println(*t.TopicArn)
	}

	return nil
}

func CreateTopic(topicName string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(topicName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(result)

	return nil
}

func ListSubscriptions() error {
	emailPtr := flag.String("e", "", "The email address of the user subscribing to the topic")
	topicPtr := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *emailPtr == "" || *topicPtr == "" {
		fmt.Println("You must supply an email address and topic ARN")
		fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
		return errors.New("must supply email and topic")
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
		return err
	}

	fmt.Println(*result.SubscriptionArn)

	return nil
}

func SubscribeToTopic(emailPtr *string, topicPtr *string) error {
	if *emailPtr == "" || *topicPtr == "" {
		fmt.Println("You must supply an email address and topic ARN")
		fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
		return errors.New("must supply email and topic")
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
		return err
	}

	fmt.Println(*result.SubscriptionArn)

	return nil
}

func SendMessageToAllTopicSubscribers(msgPtr *string, topicPtr *string) error {

	if *msgPtr == "" || *topicPtr == "" {
		fmt.Println("You must supply a message and topic ARN")
		fmt.Println("Usage: go run SnsPublish.go -m MESSAGE -t TOPIC-ARN")
		return errors.New("must supply email and topic")
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:  msgPtr,
		TopicArn: topicPtr,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(*result.MessageId)

	return nil
}
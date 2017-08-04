package main

import (
	"flag"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

var topicName = flag.String("topic", "", "topic name to subscribe to")
var subscriptionName = flag.String("subscription-name", "", "subscription name to create/read from")
var pub *pubsub.Client

func init() {
	var err error

	ctx := context.Background()
	pub, err = pubsub.NewClient(ctx, os.Getenv("GCLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error

	flag.Parse()
	defer pub.Close()

	ctx := context.Background()
	topic := pub.Topic(*topicName)

	if exists, _ := topic.Exists(ctx); !exists {
		topic, err = pub.CreateTopic(ctx, *topicName)
		if err != nil {
			log.Fatalf("Failed to create Topic: %v", err)
		}
	}

	sub := pub.Subscription(*subscriptionName)

	if exists, _ := sub.Exists(ctx); !exists {
		sub, err = pub.CreateSubscription(ctx, *subscriptionName, pubsub.SubscriptionConfig{Topic: topic})
		if err != nil {
			log.Fatalf("Failed to create Subscription: %v", err)
		}
	}

	err = sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("Got message: %s", msg.Data)
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
}

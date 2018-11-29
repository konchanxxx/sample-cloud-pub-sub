package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rexitorg/sample-cloud-pub-sub/infrastructure/pubsub"
)

func main() {
	ctx := context.Background()
	s, err := pubsub.NewPubSubService(ctx)
	if err != nil {
		log.Fatalf("Failed initialize cloud pubsub service: %#v", err)
	}

	topic, err := s.Client.CreateTopic(ctx, "konchan")
	if err != nil {
		log.Fatalf("Failed create topic: %#v", err)
	}

	fmt.Printf("Topic: %#v", topic)
}

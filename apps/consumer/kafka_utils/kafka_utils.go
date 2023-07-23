package kafka_utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alexbechmann/libs/utils/generated"
	kafka "github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func StartConsuming() {
	topic := "my-topic"
	fmt.Println("Starting consumer")

	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "consumer-group-id",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	fmt.Println("Start consuming... ")
	for {
		m, err := reader.ReadMessage(context.Background())
		handleEvent(m.Value)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func handleEvent(bytes []byte) {
	var customerEvent generated.CustomerCloudEvent
	proto.Unmarshal(bytes, &customerEvent)

	fmt.Printf("Received message: %v\n", customerEvent.Payload)
	switch customerEvent.Payload.(type) {
	case *generated.CustomerCloudEvent_Purchase:
		{
			purchase := customerEvent.GetPurchase()
			fmt.Printf("Received purchase message with id: %v\n", purchase.Id)
			break
		}

	case *generated.CustomerCloudEvent_PageView:
		{
			pageView := customerEvent.GetPageView()
			fmt.Printf("Received page view message with id: %v\n", pageView.Id)
			break
		}
	default:
		{
			break
		}
	}
}

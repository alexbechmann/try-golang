package kafka_utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"try-golang/utils/generated"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	kafka "github.com/segmentio/kafka-go"
)

func StartProducing() {
	topic := "my-topic"
	partition := 0

	fmt.Printf("Starting producer")

	brokers := os.Getenv("KAFKA_BROKERS")
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer conn.Close()

	for {
		purchase := generated.CustomerCloudEvent{
			Payload: &generated.CustomerCloudEvent_Purchase{
				Purchase: &generated.PurchaseCloudEvent{
					Id:          uuid.New().String(),
					Source:      "producer",
					SpecVersion: "1.0",
					Type:        generated.PurchaseCloudEvent_EXAMPLE_CUSTOMER_PURCHASE,
					Time:        timestamppb.Now(),
					Data: &generated.PurchaseCloudEvent_Data{
						CustomerId: "customer1",
						Amount:     1.0,
					},
				},
			},
		}

		data, err := proto.Marshal(&purchase)
		if err != nil {
			panic(err)
		}

		_, err = conn.WriteMessages(kafka.Message{Value: data})

		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		fmt.Println("Sent message", purchase.Payload)
		time.Sleep(3 * time.Second)
	}
}

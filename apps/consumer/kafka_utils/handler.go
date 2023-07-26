package kafka_utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"try-golang/consumer/store"
	"try-golang/utils/protos"

	kafka "github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type KafkaHandler interface {
	HandleCustomerEvent(event *protos.CustomerCloudEvent);
	StartConsuming()
}

type KafkaHandlerImpl struct {
	store *store.Store
}

func (handler KafkaHandlerImpl) HandleCustomerEvent(event *protos.CustomerCloudEvent) {
	fmt.Println("Received message: ", event.Payload)
	fmt.Println("Store: ", handler.store)
	fmt.Printf("Received message: %v\n", event.Payload)
	switch event.Payload.(type) {
	case *protos.CustomerCloudEvent_Purchase:
		{
			purchase := event.GetPurchase()
			fmt.Printf("Received purchase message with id: %v\n", purchase.Id)
			break
		}

	case *protos.CustomerCloudEvent_PageView:
		{
			pageView := event.GetPageView()
			fmt.Printf("Received page view message with id: %v\n", pageView.Id)
			break
		}
	default:
		{
			break
		}
	}
}

func (handler KafkaHandlerImpl) TestConnection() {
	fmt.Println("Test connection")
}

func (handler KafkaHandlerImpl) StartConsuming() {
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
		if err != nil {
			log.Fatalln(err)
		}
		var customerEvent protos.CustomerCloudEvent
		proto.Unmarshal(m.Value, &customerEvent)
		handler.HandleCustomerEvent(&customerEvent)
	}

}
func KakfaHandlerProvider(store store.Store) KafkaHandler {
	return &KafkaHandlerImpl{
		store: &store,
	}
}
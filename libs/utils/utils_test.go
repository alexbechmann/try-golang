package utils

import (
	"fmt"
	"testing"

	"try-golang/utils/protos"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDoSomething(t *testing.T) {
	fmt.Println("Something")

	purchase := protos.CustomerCloudEvent{
		Payload: &protos.CustomerCloudEvent_Purchase{
			Purchase: &protos.PurchaseCloudEvent{
				Id:          "id1",
				Source:      "lib",
				SpecVersion: "1.0",
				Type:        protos.PurchaseCloudEvent_EXAMPLE_CUSTOMER_PURCHASE,
				Time:        timestamppb.Now(),
				Data: &protos.PurchaseCloudEvent_Data{
					CustomerId: "customer1",
					Amount:     1.0,
				},
			},
		},
	}

	fmt.Println(purchase.Payload)

	customerId := purchase.GetPurchase().Data.CustomerId

	if customerId != "customer1" {
		t.Error("Expected customer1 ..., got ", customerId)
	}
}

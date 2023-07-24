package utils

import (
	"fmt"

	"try-golang/utils/protos"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func DoSomething() {
	fmt.Println("Do Something")

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
}

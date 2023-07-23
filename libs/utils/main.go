package utils

import (
	"fmt"

	"github.com/alexbechmann/libs/utils/generated"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Something() {
	fmt.Println("Something")

	purchase := generated.CustomerCloudEvent{
		Payload: &generated.CustomerCloudEvent_Purchase{
			Purchase: &generated.PurchaseCloudEvent{
				Id:          "id1",
				Source:      "lib",
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

	fmt.Println(purchase.Payload)
}

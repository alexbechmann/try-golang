package kafkahandler

import (
	"testing"

	"try-golang/consumer/testutils"
	"try-golang/utils/protos"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestKafkaHandler_Purchase_Event(t *testing.T) {
	mock_controller := gomock.NewController(t)
	store := testutils.NewMockStore(mock_controller)
	store.EXPECT().SavePurchase(gomock.Any()).Return(true).Times(1)

	kafkaHandler := &KafkaHandlerImpl{
		store: store,
	}

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

	kafkaHandler.HandleCustomerEvent(&purchase)
	assert.Equal(t, mock_controller.Satisfied(), true)
}

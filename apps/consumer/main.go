package main

import (
	"fmt"
	"time"

	"try-golang/consumer/kafka_utils"
	"try-golang/consumer/store"
	"try-golang/utils"

	"go.uber.org/fx"
)
func main() {
	items := [2]string{
		"a", "b",
	}

	for index, item := range items {
		fmt.Println(fmt.Sprint(index), ": ", item)
	}

	fmt.Println("The time is: ", time.Now().Round(0))

	utils.DoSomething()

	fx.New(
		fx.Provide(store.StoreProvider),
		fx.Provide(kafka_utils.KakfaHandlerProvider),
		fx.Invoke(func (store store.Store, handler kafka_utils.KafkaHandler) {
			store.Connect()
			handler.StartConsuming()
		}),
	).Run()
}

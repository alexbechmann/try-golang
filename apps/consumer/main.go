package main

import (
	"fmt"
	"time"

	"try-golang/consumer/kafkahandler"
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
		fx.Provide(kafkahandler.KakfaHandlerProvider),
		fx.Invoke(func(store store.Store, handler kafkahandler.KafkaHandler) {
			store.Connect()
			handler.StartConsuming()
		}),
	).Run()
}

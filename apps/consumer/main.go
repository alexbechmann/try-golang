package main

import (
	"fmt"
	"time"

	"github.com/alexbechmann/try-golang/apps/consumer/kafka_utils"
	"github.com/alexbechmann/try-golang/libs/utils"
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

	kafka_utils.StartConsuming()
}

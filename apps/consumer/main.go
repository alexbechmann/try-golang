package main

import (
	"fmt"
	"time"

	"try-golang/consumer/kafka_utils"
	"try-golang/utils"
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

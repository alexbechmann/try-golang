package main

import (
	"fmt"
	"time"

	"try-golang/producer/do"
	"try-golang/producer/kafka_utils"
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

	do.DoSomething2()
	utils.DoSomething()

	kafka_utils.StartProducing()
}

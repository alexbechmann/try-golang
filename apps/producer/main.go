package main

import (
	"fmt"
	"time"

	"github.com/alexbechmann/try-golang/apps/producer/do"
	"github.com/alexbechmann/try-golang/apps/producer/kafka_utils"
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

	kafka_utils.StartProducing()
}

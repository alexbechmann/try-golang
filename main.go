package main

import (
	"fmt"
	"time"

	"try-go/do"
)

func main() {
	// array of names
	items := [2]string{
		"a", "b",
	}

	for index, item := range items {
		message := fmt.Sprintf(fmt.Sprint(index), ": ", item)
		fmt.Println(message)
	}

	do.DoSomething()
	fmt.Println("The time is: ", time.Now().Round(0))
}

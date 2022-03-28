package main

import (
	"fmt"

	"go-playground/do"
)

func main() {
	// array of names
	items := [2]string{
		"a", "b",
	}

	for index, item := range items {
		message := fmt.Sprintf("%d: %s", index, item)
		fmt.Println(message)
	}

	do.DoSomething()
}

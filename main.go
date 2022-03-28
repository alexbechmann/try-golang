package main

import (
	"do"
	"fmt"
)

// b "basic/basic"
// b "myapp/utils/do"

// "github.com/fatih/color"

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

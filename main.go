package main

import (
	"fmt"
	"time"

	"try-go/do"
	"try-go/geo"
)

func main() {
	items := [2]string{
		"a", "b",
	}

	for index, item := range items {
		fmt.Println(fmt.Sprint(index), ": ", item)
	}

	do.DoSomething()
	fmt.Println("The time is: ", time.Now().Round(0))

	grid := geo.ExampleFromGeo(geo.Args{
		Lat: 37.775938728915946,
		Lng: -122.41795063018799,
	})
	fmt.Println("The grid is: ", grid)
}

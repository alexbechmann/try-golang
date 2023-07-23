package geo

import (
	"fmt"

	"github.com/uber/h3-go/v4"
)

type Args struct {
	Lat float64
	Lng float64
}

func ExampleFromGeo(args Args) string {
	latLng := h3.NewLatLng(args.Lat, args.Lng)
	resolution := 9 // between 0 (biggest cell) and 15 (smallest cell)

	fmt.Print("Using lat ", latLng.Lat)
	fmt.Println(" and lng ", latLng.Lng)

	cell := h3.LatLngToCell(latLng, resolution)
	fmt.Printf("%#x\n", cell)
	// Output:
	// 0x8928308280fffff
	return cell.String()
}

package geo

import (
	"fmt"

	"github.com/uber/h3-go/v4"
)

func ExampleFromGeo() string {
	latLng := h3.NewLatLng(37.775938728915946, -122.41795063018799)
	resolution := 9 // between 0 (biggest cell) and 15 (smallest cell)

	cell := h3.LatLngToCell(latLng, resolution)
	fmt.Printf("%#x\n", cell)
	// Output:
	// 0x8928308280fffff
	return cell.String()
}

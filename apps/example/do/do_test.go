package do

import "testing"

func TestDoSomething(t *testing.T) {
	result := DoSomething()

	if result != "Do something..." {
		t.Error("Expected Do something..., got ", result)
	}
}

package do

import "testing"

func TestDoSomething2(t *testing.T) {
	result := DoSomething2()

	if result != "Do something..." {
		t.Error("Expected Do something..., got ", result)
	}
}

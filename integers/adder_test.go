package integers

import (
	"fmt"
	"testing"
)

// Output comment to test the output of the example function
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d", expected, sum)
	}
}

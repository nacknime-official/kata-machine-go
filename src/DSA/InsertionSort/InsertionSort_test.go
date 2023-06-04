package insertionsort

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	InsertionSort(arr)

	expected := []int{3, 4, 7, 9, 42, 69, 420}

	if !dsa.ArraysEqual(expected, arr) {
		t.Errorf("Expected result: %v, got: %v", expected, arr)
	}
}

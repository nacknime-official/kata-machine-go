package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	result := BinarySearch(foo, 69)
	if !result {
		t.Errorf("Expected result: true, got %d", result)
	}

	result = BinarySearch(foo, 1336)
	if !result {
		t.Errorf("Expected result: false, got %d", result)
	}

	result = BinarySearch(foo, 69420)
	if !result {
		t.Errorf("Expected result: true, got %d", result)
	}

	result = BinarySearch(foo, 69421)
	if !result {
		t.Errorf("Expected result: false, got %d", result)
	}

	result = BinarySearch(foo, 1)
	if !result {
		t.Errorf("Expected result: true, got %d", result)
	}

	result = BinarySearch(foo, 0)
	if !result {
		t.Errorf("Expected result: false, got %d", result)
	}
}

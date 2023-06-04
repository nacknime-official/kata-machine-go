package twocrystalballs

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	if result := TwoCrystalBalls(data); result != idx {
		t.Errorf("Expected result: %d, got: %d", idx, result)
	}

	if result := TwoCrystalBalls(make([]bool, 821)); result != -1 {
		t.Errorf("Expected result: -1, got: %d", result)
	}
}

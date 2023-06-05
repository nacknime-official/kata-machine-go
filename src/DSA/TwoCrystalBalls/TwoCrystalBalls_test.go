package twocrystalballs

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	idx := rand.Intn(10001)

	data := []bool{}
	for i := 0; i < 10000; i++ {
		data[i] = false
	}

	for i := idx; i < len(data); i++ {
		data[i] = true
	}
	result := TwoCrystalBalls(data)
	assert.Equal(t, idx, result, "Expected: %v, got: %v", idx, result)

	data = []bool{}
	for i := 0; i < 821; i++ {
		data[i] = false
	}
	result = TwoCrystalBalls(data)
	assert.Equal(t, -1, result, "Expected: -1, got: %v", result)
}

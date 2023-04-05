package arraylist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestArrayList(t *testing.T) {
	list := NewArrayList[int](3)
	dsa.TestList(t, list)
}

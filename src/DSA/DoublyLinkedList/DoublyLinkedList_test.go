package doublylinkedlist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	dsa.TestList(t, list)
}

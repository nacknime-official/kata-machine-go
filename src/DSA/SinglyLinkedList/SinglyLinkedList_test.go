package singlylinkedlist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	dsa.TestList(t, list)
}

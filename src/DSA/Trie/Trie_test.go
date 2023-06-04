package trie

import (
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foo")
	trie.Insert("fool")
	trie.Insert("foolish")
	trie.Insert("bar")

	result := trie.Find("fo")
	sort.Strings(result)

	expected := []string{
		"foo",
		"fool",
		"foolish",
	}

	if !stringSlicesEqual(result, expected) {
		t.Errorf("Expected result: %v, got: %v", expected, result)
	}

	expected = []string{
		"foo",
		"foolish",
	}

	trie.Delete("fool")
	result = trie.Find("fo")

	if !stringSlicesEqual(result, expected) {
		t.Errorf("Expected result: %v, got: %v", expected, result)
	}
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

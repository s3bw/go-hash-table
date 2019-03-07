package hashtable

import (
	"testing"

	"github.com/foxyblue/a-hash-table/hashtable"
)

var hashTests = []struct {
	s        string // String to hash
	p        int    // Prime number
	b        int    // Number of buckets
	expected int    // Expected outcome
}{
	{"cat", 151, 53, 5},
	{"paul", 151, 53, 29},
}

// TestHash deduces the consistency of our Hash function
func TestHash(t *testing.T) {
	for _, tt := range hashTests {
		result := hashtable.Hash(tt.s, tt.p, tt.b)
		if result != tt.expected {
			t.Errorf("Expected: %d, got: %d", tt.expected, result)
		}
	}
}

var hashFunctionTests = []struct {
	s        string // String to hash
	p        int    // Prime number
	a        int    // Hash Attempt
	expected int    // Expected outcome
}{
	{"paul", 53, 0, 29},
	{"paul", 53, 1, 13},
}

func TestHashFunction(t *testing.T) {
	for _, tt := range hashFunctionTests {
		result := hashtable.HashFunction(tt.s, tt.p, tt.a)
		if result != tt.expected {
			t.Errorf("Expected: %d, got: %d", tt.expected, result)
		}
	}
}

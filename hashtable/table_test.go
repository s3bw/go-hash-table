package hashtable

import (
	"bytes"
	"testing"

	"github.com/foxyblue/a-hash-table/hashtable"
)

func TestHashTableInsert(t *testing.T) {
	// Key to insert
	key := "cat"
	// Value to insert
	value := []byte("cat")

	hashTable := hashtable.NewHashTable()
	if hashTable.Count != 0 {
		t.Errorf("Expected: 0, got: %d", hashTable.Count)
	}
	hashTable.Insert(key, value)
	if hashTable.Count != 1 {
		t.Errorf("Expected: 1, got: %d", hashTable.Count)
	}
}

var hashTableSearchTest = []struct {
	search        string // Key to search for
	expectedValue []byte // Value expected
	expectToFind  bool   // If the test should find an item
}{
	// Case for finding something that exists
	{"cat", []byte("ludwig"), true},
	// Case for finding something that does not exist
	{"no cat", nil, false},
}

func TestHashTableSearch(t *testing.T) {
	hashTable := hashtable.NewHashTable()
	hashTable.Insert("cat", []byte("ludwig"))
	for _, tt := range hashTableSearchTest {
		result := hashTable.Search(tt.search)
		if tt.expectToFind {
			// bytes.Compare return 0 if the []byte are the same
			// we need to raise the error if they are not the same.
			if !(bytes.Compare(result.Value, tt.expectedValue) == 0) {
				t.Errorf("Expected: %s, got: %s", tt.expectedValue, result.Value)
			}
		} else {
			if result != nil {
				t.Errorf("Did not expect to find something!")
			}
		}
	}
}

func TestHashTableDelete(t *testing.T) {
	key := "cat"
	value := []byte("ludwig")

	hashTable := hashtable.NewHashTable()
	hashTable.Insert(key, value)
	if hashTable.Count != 1 {
		t.Errorf("Expected: 1, got: %d", hashTable.Count)
	}
	hashTable.Delete(key)
	if hashTable.Count != 0 {
		t.Errorf("Expected: 0, got: %d", hashTable.Count)
	}
	result := hashTable.Search(key)
	if result != nil {
		t.Error("Expected: <nil>, got a (type *Item) instead.")
	}
}

var resizeHashTableTest = []struct {
	expectedSize  int  // Expected size of the hash table
	expectedCount int  // Expected count of the hash table
	resizeUp      bool // If true resize up else resize down
}{
	{hashtable.BaseHashTableSize, 1, true},
	{107, 1, false},
	{hashtable.BaseHashTableSize, 1, true},
}

func TestHashTableResize(t *testing.T) {
	key := "cat"
	value := []byte("ludwig")
	hashTable := hashtable.NewHashTable()
	hashTable.Insert(key, value)
	for _, tt := range resizeHashTableTest {
		if hashTable.Size != tt.expectedSize {
			t.Errorf("Expected: %d, got: %d", tt.expectedSize, hashTable.Size)
		} else if hashTable.Count != tt.expectedCount {
			t.Errorf("Expected: %d, got: %d", tt.expectedCount, hashTable.Count)
		}
		if tt.resizeUp {
			hashTable.ResizeUp()
		} else {
			hashTable.ResizeDown()
		}
	}

}

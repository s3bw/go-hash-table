package hashtable

import (
	"testing"

	"github.com/foxyblue/a-hash-table/hashtable"
)

func TestNextPrime(t *testing.T) {
	result := hashtable.NextPrime(54)
	if result != 59 {
		t.Errorf("Expected: 59, got: %d", result)
	}
}

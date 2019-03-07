/* This module implements the hash table structure.

Since Go is garbage collected we can avoid freeing
memory (as is implemented in the C tutorial). */
package hashtable

type Dictionary interface {
	// Insert consider resize or returning err if full
	Insert(key string, value []byte)
	Search(key string) *Item
	Delete(key string) error
}

// Item storing the key-value pairs in a struct
type Item struct {
	Key string

	// Value will be a slice of bytes, this will make
	// the type more flexible
	Value []byte

	// Deleted represents if the item was deleted since
	// Go does not support const structs
	Deleted bool
}

type HashTable struct {
	Size  int
	Count int
	Items []*Item
}

func newItem(k string, v []byte) *Item {
	return &Item{
		Key:     k,
		Value:   v,
		Deleted: false,
	}
}

func NewHashTable() *HashTable {
	size := 53
	return &HashTable{
		Size:  size,
		Count: 0,
		Items: make([]*Item, size),
	}
}

func (ht *HashTable) Insert(key string, value []byte) {
	item := newItem(key, value)
	for a := 0; a < ht.Size; a++ {
		index := HashFunction(item.Key, ht.Size, a)
		if ht.Items[index] == nil || ht.Items[index].Deleted {
			ht.Items[index] = item
			ht.Count++
			break
		} else if ht.Items[index] != nil {
			// Update the item if the insert happens again
			// Count does not need to be incremented since
			// we are just updating an item's Value.
			if ht.Items[index].Key == key {
				ht.Items[index] = item
				break
			}
		}
	}
}

func (ht *HashTable) Search(key string) *Item {
	for a := 0; a < ht.Size; a++ {
		index := HashFunction(key, ht.Size, a)
		item := ht.Items[index]
		if item == nil {
			return nil
		} else {
			if item.Key == key && !item.Deleted {
				return item
			}
		}
	}
	return nil
}

func (ht *HashTable) Delete(key string) {
	for a := 0; a < ht.Size; a++ {
		index := HashFunction(key, ht.Size, a)
		item := ht.Items[index]
		if item == nil {
			break
		} else {
			if !item.Deleted {
				if item.Key == key {
					ht.Items[index].Deleted = true
					ht.Count--
					break
				}
			}
		}
	}
}

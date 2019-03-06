/* This module implements the hash table structure.

Since Go is garbage collected we can avoid freeing
memory (as is implemented in the C tutorial). */
package main

// Item storing the key-value pairs in a struct
type Item struct {
	key   *rune
	value *rune
}

type HashTable struct {
	Size  int
	Count int
	// To use a slice or a double pointer?
	Items []*Item
}

func newItem(k, v *rune) *Item {
	return &Item{
		key:   k,
		value: v,
	}
}

func newHashTable() *HashTable {
	size := 53

	return &HashTable{
		Size:  size,
		Count: 0,
		Items: make([]*Item, size),
	}
}

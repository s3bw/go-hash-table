package main

import (
	"fmt"

	"github.com/foxyblue/a-hash-table/hashtable"
)

func main() {
	// Test this is 5
	fmt.Println(hashtable.Hash("cat", 151, 53))

	cat := "paul"

	// Test this is 29
	result := hashtable.Hash(cat, 151, 53)
	fmt.Println(result)

	// Test this is 29
	result1 := hashtable.HashFunction(cat, 53, 0)
	fmt.Println(result1)

	// Test this is 13
	result2 := hashtable.HashFunction(cat, 53, 1)
	fmt.Println(result2)

	catValue := []byte("cat")
	hashTable := hashtable.NewHashTable()
	// Test this is 0
	fmt.Println(hashTable.Count)
	hashTable.Insert(cat, catValue)
	// Test this is 1
	fmt.Println(hashTable.Count)

	// Test expect cat
	search1 := hashTable.Search(cat)
	if search1 != nil {
		fmt.Println(string(search1.Value))
	}

	// Test expect nil
	search2 := hashTable.Search("no cat")
	if search2 == nil {
		fmt.Println("Did not find anything.")
	}

	hashTable.Delete(cat)
	// Test this is 0
	fmt.Println(hashTable.Count)
	search3 := hashTable.Search(cat)
	if search3 == nil {
		// Test expect nil
		fmt.Println("Did not find anything.")
	}

	// Test expect 59
	prime := hashtable.NextPrime(54)
	fmt.Println(prime)

	// Test base size goes from 51 -> 102 -> 51
	// Count remains 1
	hashTable.Insert(cat, catValue)
	fmt.Println(hashTable.BaseSize)
	fmt.Println(hashTable.Count)

	hashTable.ResizeUp()
	fmt.Println(hashTable.BaseSize)
	fmt.Println(hashTable.Count)

	hashTable.ResizeDown()
	fmt.Println(hashTable.BaseSize)
	fmt.Println(hashTable.Count)

}

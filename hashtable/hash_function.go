package hashtable

import (
	"math"
)

const (
	PrimeOne = 151
	PrimeTwo = 17
)

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Hash(s string, prime, mod int) int {
	hash := 0
	len_s := len(s)
	for i := 0; i < len_s; i++ {
		hash += pow(prime, len_s-(i+1)) * int(s[i])
		hash = hash % mod
	}
	return int(hash)
}

// HashFunction uses double hashing to deal with collisions
func HashFunction(s string, buckets, attempt int) int {
	hashA := Hash(s, PrimeOne, buckets)
	hashB := Hash(s, PrimeTwo, buckets)
	return (hashA + (attempt * (hashB + 1))) % buckets
}

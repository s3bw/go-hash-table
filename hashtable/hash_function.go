package main

import (
	"fmt"
	"math"
)

func pow(a, b float64) int {
	return int(math.Pow(a, b))
}

func Hash(s []byte, a, m float64) int {
	hash := 0
	len_s := len(s)
	for i := 0; i < len_s; i++ {
		hash += pow(a, float64(len_s-(i+1))) * int(s[i])
		hash = hash % int(m)
	}
	return int(hash)
}

func main() {
	cat := []byte("cat")
	result := Hash(cat, 151, 53)
	fmt.Println(result)
}

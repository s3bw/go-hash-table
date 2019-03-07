/* Our array size should be a prime number roughly half or
 * double the current size. Finding new array size isn't
 * trivial. To do so, we store a base size, which we want
 * the array to be, and then define the actual size as the
 * first prime larger than the base size. To resize up, we
 * double the base size, and find the first larger prime,
 * and to resize down, we halve the size and find the next
 * larger prime.
 */

package hashtable

import "math"

// IsPrime determines whether x is prime or not
// 1 - prime
// 0 - not prime
// -1 - undefine (i.e. x < 2)
func IsPrime(x int) int {
	if x < 2 {
		return -1
	}
	if x < 4 {
		return 1
	}
	if (x % 2) == 0 {
		return 0
	}
	for i := 3; i <= int(math.Sqrt(float64(x))); i += 2 {
		if (x % i) == 0 {
			return 0
		}
	}
	return 1
}

// NextPrime returns the next prime number after x or x if x
// is prime
func NextPrime(x int) int {
	for IsPrime(x) != 1 {
		x++
	}
	return x
}

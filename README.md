# Implementing a hash table in go

As per implementation in C, Source: https://github.com/jamesroutley/write-a-hash-table/

- Go map Source: https://github.com/golang/go/blob/master/src/runtime/map.go

## To-Do

- Step 06
- Appendix
- Write tests
- Implement as redis server.

## Interface

Associative arrays are a collection of unordered key-value pairs. Duplicate keys are not
permitted.

The following operations are supported:

- `search(a, k)`: return the value `v` associated with key `k` from the associative array
  `a`, or `NULL` if the key does not exist.
- `insert(a, k, v)`: store the pair `k:v` in the associative array `a`.
- `delete(a, k)`: delete the `k:v` pair associated with `k`, or do nothing if `k` does not
  exist.

## Terminology

- Associative array: an abstract data structure which implements the interface described
  above. Also known as a map, symbol table or dictionary.
- Hash table: a fast implementation of the associative array interface which makes use of
  a hash function. Also called a hash map, map, hash or dictionary.


## Hash Functions

- Take a string and return a number between 0 and `m`. `m` being our desired bucket array
  length.
- Return evenly distributed bucket indexes for an average set of inputs. Unevenly distributed
  hash function will put more items in some buckets that others. Leading to high rates of
  collision.

As a pseudocode example:

```python
def hash(string, a, num_buckets):
    hash = 0
    string_len = len(string)
    for i in range(string_len):
        hash += (a ** (string_len - (i+1))) * char(string[i])
    hash = hash % num_buckets
    return hash


>>>hash("cat", 151, 53) -> ((151**2 * 99) + (151**1 * 97) + (151**0 * 116)) % 53
5
```

Changing the value of `a` gives us a different hash function. `hash("cat, 163, 53) = 3`

### Handling collisions

Different input keys will map to the same array index, this causes a collision. Hash
tables need to implement a method of dealing with these collisions.

The tutorial uses a technique called open addressing with double hashing. Double hashing
makes use of two hash functions to calculate the index an item should be stored at after
`i` collisions.

#### Double Hashing

The index that should be used after `i` collisions is given by

```
index = hash_a(string) + i * hash_b(string) % num_buckets
```

Given the function above `0` collisions will indicate the bucket is determined just by
`hash_a`

#### Deleting

When using open addressing, deleting is slightly more complicated then inserting or
searching. The item we wish to delete may be part of a collision chain. Removing it would
break that chain and finding the items in the tail of the chain impossible. To solve this,
we simply mark it as deleted.

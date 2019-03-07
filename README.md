# Implementing a hash table in go

As per implementation in C, Source: https://github.com/jamesroutley/write-a-hash-table/

- Go map Source: https://github.com/golang/go/blob/master/src/runtime/map.go

## To-Do

- Implement as redis server.

## Testing

```
go test hashtable/*
```

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

## Resizing

As more items are inserted into a hash table the chance of collision increases and the
hash table will eventually fail if we don't resize it to be able to include more items.

A hash table's `load` is the ratio of filled buckets to total buckets.

We will aim to resize up when `load > 0.7` and resize down when `load < 0.1`.

To resize we create a new hash table roughly half or twice as big as the current, and
insert all non-deleted items.


## Alternative approaches to collision handling

 - Separate chaining
 - Open addressing

### Separate chaining

Under separate chaining. each bucket contains a linked list. When items collide, they are
added to the list. Methods:

- `Insert`: hash the key to get the bucket index. If there is nothing in that bucket,
  store the item there. If there is already an item there, append the item to the linked
  list.
- `Search`: hash the key to get the bucket index. Traverse the linked list, comparing each
  item's key to the search key. If the key is found, return the value, else return `nil`.
- `Delete` hash the key to get the bucket index. Traverse the linked list, comparing each
  item's key to the delete key. If the key is found, remove the item from the linked list.
  If there is only one item in the linked list, place the `nil` pointer in the bucket, to
  indicate that it is empty.

This has the advantage of being simple to implement, but is space inefficient. Each item
has to also store a pointer to the next item in the linked list, or the `nil` pointer if
no items come after it. This is space wasted on bookkeeping, which could be spent on
storing more items.

### Open addressing

Open addressing aims to solve the space inefficiency of separate chaining. When
collisions happen, the collided item is place in some other bucket in the table. The
bucket that the item is placed into is chosen according to some predetermined rule, which
can be repeated when searching for the item. There are three common methods for choosing
the bucket to insert a collided item into.

Note: Open addressing can lead to more concise tables with better cache performance than
bucketing, but performance will be more brittle as the load factor (ratio of occupancy to
capacity) of the hash table starts to get high.

#### Linear probing

When a collision occurs, the index is incremented and the item is put in the next
available bucket in the array.

Linear probing offers good cache performance, but suffers from clustering issues. Putting
collided items in the next available bucket can lead to long contiguous stretches of
filled buckets, which need to be iterated over when inserting, searching or deleting.

#### Quadratic probing

Similar to linear probing, but instead of putting the collided item in the next available
bucket, we try to put it in the buckets whose indexes follow the sequence: `i, i + 1, i +
4, i + 9, i + 16, ... `, where `i` is the original hash of the key.

Quadratic probing reduces, but does not remove, clustering, and still offers decent cache
performance.

#### Double hashing

Double hashing aims to solve the clustering problem. To do so, we use a second hash
function to choose a new index for the item. Using a hash function gives us a new bucket,
the index of which should be evenly distributed across all buckets. This removes
clustering, but also removes any boosted cache performance from locality of reference.
Double hashing is a common method of collision management in production hash tables.

## How big should a hash table be?

The hash function maps keys to an integer between `0 -> m -1`.

With bucketing, `m` should be about the same as the maximum number of items you expect to
put into the table. With open addressing, make it (say) 30% larger or more. Selecting `m`
to be a prime number minimizes the danger of a bad hash function.


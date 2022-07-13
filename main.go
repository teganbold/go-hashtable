package main

import "fmt"

const ArraySize = 7

// Hash Table Structure
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket Structure (linked list)
type bucket struct {
	head *bucketNode
}

// bucket node strucutre
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take a key and return true if in the table and false if not
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from ther hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

// Search will take in a key and return true if the butcket exists
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// insert will tke in a key, create a node with the key and inside the node into the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Print(k, " already exists")
	}

}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

// Define Hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"HOMER",
		"MARGE",
		"BART",
		"LISA",
		"MAGGIE",
		"MOE",
		"NED",
		"SEYMOUR",
		"PATTY",
		"SELMA",
		"WAYLON",
		"MONTGOMERY",
		"ABRAHAM",
		"MILHOUSE",
		"BARNEY",
		"LENNY",
		"CARL",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("NED")
}

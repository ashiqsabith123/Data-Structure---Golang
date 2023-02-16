package main

import "fmt"

const ArraySize = 10

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

func (b *bucket) insert(key string) {
	newNode := new(bucketNode)
	newNode.key = key

	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(key string) bool {
	temp := b.head

	for temp != nil {
		if temp.key == key {
			return true
		}
		temp = temp.next
	}

	return false
}

func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}
	temp := b.head

	for temp.next != nil {
		if temp.next.key == k {
			temp.next = temp.next.next
			return
		}
	}

	fmt.Println(k, "Doens`t exist....!")
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

func main() {
	table := HashTable{}
	table.Insert("hello")

}

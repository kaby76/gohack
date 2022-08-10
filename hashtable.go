/*
Package hashtable provides a Hash Table
*/
package antlr

import (
//	"math"
	"fmt"
)

const arrayLength = 1000

// HashTable : a hash table data structure
type HashTable struct {
	data [arrayLength]*LinkedList
	count int
}

type listData struct {
	key   *DFAState
	value *DFAState
}

// New : a hash table constructor
func NewHashTable() *HashTable {
	return &HashTable{
		[arrayLength]*LinkedList{},
		0,
	}
}

func (h *HashTable) MyIterator() <-chan *DFAState {
	c := make(chan *DFAState)
	go func() {
		defer close(c)
		for x := range h.data {
			if h.data[x] == nil {
				continue
			}
			node := h.data[x].Head
			for {
				if node != nil {
					d := node.Data.(listData)
					c <- d.key
					break
				} else {
					break
				}
				node = node.Next
			}
		}
	}()
	return c
}

func hash(s *DFAState) int {
//	h := 0
	// char[0] * 31^n-1 + char[1] * 31^n-2 + ... + char[n-1]
//	for pos, char := range s {
//		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
//	}
	if false {
		fmt.Print("hash ")
		fmt.Print(s)
		fmt.Print(" => ")
	}
	hash := s.hash()
	if false {
		fmt.Println(hash)
	}
	return hash
}

func equals(one *DFAState, other *DFAState) bool {
	if false {
		fmt.Print("equals ")
		fmt.Print(one)
		fmt.Print(" ")
		fmt.Print(other)
		fmt.Print(" => ")
	}
	result := one.equals(other)
	if false {
		fmt.Println(result)
	}
	return result
}

func index(hash int) int {
	return hash % arrayLength
}

// Set : set a key and value
func (h *HashTable) Set(k *DFAState, v *DFAState) *HashTable {
	if false {
		fmt.Print("Set ")
		fmt.Print(k)
		fmt.Print(" ")
		fmt.Println(v)
	}
	index := index(hash(k))
	if false {
		fmt.Print("index ")
		fmt.Println(index)
	}
	if h.data[index] == nil {
		h.data[index] = NewLinkedList()
		h.data[index].Append(listData{k, v})
		h.count++
	} else {
		node := h.data[index].Head
		for {
			if node != nil {
				d := node.Data.(listData)
				if equals(k, d.key) {
					d.value = v
					if false {
						fmt.Print("Entered ")
						fmt.Print(d.key)
						fmt.Print(" ")
						fmt.Println(v)
					}
					break
				}
			} else {
				h.data[index].Append(listData{k, v})
				h.count++
				break
			}
			node = node.Next
		}
	}
	return h
}

func (h *HashTable) Len() int {
	return h.count
}

// Get : get a value by key
func (h *HashTable) Get(k *DFAState) (result *DFAState, ok bool) {
	if false {
		fmt.Print("Get ")
		fmt.Println(k)
	}
	index := index(hash(k))
	if false {
		fmt.Print("index ")
		fmt.Println(index)
	}
	linkedList := h.data[index]

	if linkedList == nil {
		return nil, false
	}
	node := linkedList.Head
	for {
		if node != nil {
			d := node.Data.(listData)
			if equals(d.key, k) {
				return d.value, true
			}
		} else {
			return nil, false
		}
		node = node.Next
	}
}

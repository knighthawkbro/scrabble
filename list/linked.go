package list

import (
	"fmt"
	"math/rand"
	"time"
)

// Node - Defines the structure for each individual node in a linked list
type Node struct {
	// Pointer to the next Node
	next *Node
	// Pointer to the list it is attached to
	list *List
	// Value of Node
	val interface{}
}

// Next returns the next node in the list
func (node Node) Next() *Node {
	// returns nil if there is not list AND if the pointer to the next
	// node is the same as the head's next node there for there is next node
	if next := node.next; node.list != nil && next != &node.list.head {
		return next
	}
	return nil
}

// Front returns the first element of list l or nil.
func (l *List) Front() *Node {
	if l.size == 0 {
		return nil
	}
	return l.head.next
}

// GetVal returns the value stored in the specific Node
func (node *Node) GetVal() interface{} { return node.val }

func (node *Node) String() string {
	return fmt.Sprintf("[ %v ]", node.val)
}

// List - The container for all the linked nodes in a set
type List struct {
	head Node // the begining node
	size int  // size of the list
}

// Init Generates a linked list with Size=0 and head pointing to itself
func (l *List) Init() *List {
	l.head.next = &l.head
	l.size = 0
	return l
}

// New returns an initialized list.
func New() *List { return new(List).Init() }

// Len returns the length variable for the list as an integer
func (l *List) Len() int { return l.size }

// Add returns the node in a singly linked list, just adds to the front of the list
func (l *List) Add(v interface{}) error {
	if v == nil {
		return fmt.Errorf("cannot add nil value")
	}
	new := &Node{val: v, list: l}
	prev := l.head.next
	l.head.next = new
	new.next = prev
	l.size++
	return nil
}

// Remove removes the first item on a list and returns it
func (l *List) Remove() interface{} {
	if l.size == 0 {
		return nil
	}
	result := &l.head
	l.head = *l.head.next
	l.size--
	return result
}

// Get returns the first item list
func (l *List) Get() interface{} {
	if l.size == 0 {
		return nil
	}
	return l.head.next.val
}

// Contains returns true or false whether an item was contained in the list
func (l *List) Contains(i interface{}) bool {
	for current := l.Front(); current != nil; current = current.Next() {
		if current.val == i {
			return true
		}
	}
	return false
}

// Size returns the size of the list
func (l *List) Size() int {
	return l.size
}

// String allows for the fmt.Print* functions to print the list struct as a string.
func (l *List) String() string {
	if l.size == 0 {
		return "[ ]"
	}
	result := "[ "
	for current := l.Front(); current != nil; current = current.Next() {
		result += fmt.Sprintf("%v ", current.val)
	}
	return result + "]"
}

// RemoveItem takes a single value and returns true or false if that item was removed.
func (l *List) RemoveItem(i interface{}) bool {
	prev := l.Front()
	for current := l.Front(); current != nil; current = current.Next() {
		if current.val == i && current == l.Front() {
			l.Remove()
			return true
		} else if current.val == i {
			prev.next = current.next
			l.size--
			return true
		}
		prev = current
	}
	return false
}

// RemoveRandom grabs a pseudo-random number using the unix time and removes an item at
// the generated location.
func (l *List) RemoveRandom() interface{} {
	if l.size == 0 {
		return nil
	}
	rand.Seed(time.Now().UTC().UnixNano())
	// pseudo random number based on time.
	random := rand.Intn(l.size)
	rm := l.Front()
	for i := 0; i <= random; i++ {
		rm = rm.next
	}
	l.RemoveItem(rm.val)
	return rm.val
}

// GetRandom returns an item pseudo-randomly using unix time
func (l *List) GetRandom() interface{} {
	if l.size == 0 {
		return nil
	}
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(l.size)
	current := l.Front()
	for i := 0; i < random; i++ {
		current = current.next
	}
	return current
}

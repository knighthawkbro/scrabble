package array

import (
	"fmt"
	"math/rand"
	"time"
)

// Array (Public) - Structure for the Array Datatype
type Array struct {
	size       int
	collection []interface{}
}

// Init (Public) - Initializes the array with the size provided. Can be overridden by the user.
func (a *Array) Init(capacity int) *Array {
	if capacity < 0 {
		return nil
	}
	a.collection = make([]interface{}, capacity)
	a.size = 0
	return a
}

// New (Public) - Returns an initialized array with default size.
func New() *Array { return new(Array).Init(10) }

// Add (Public) - Returns an error if adding an item to the end of the array fails or not
func (a *Array) Add(item interface{}) error {
	if item == nil {
		return fmt.Errorf("cannot add nil value")
	}
	a.ensureSpace()
	a.collection[a.size] = item
	a.size++
	return nil
}

// Remove (Public) - removes the last item in the array
func (a *Array) Remove() interface{} {
	if a.size == 0 {
		return nil
	}
	removed := a.collection[a.size-1]
	a.collection[a.size-1] = nil
	a.size--
	return removed
}

// Get (Public) - Returns the last item in the array
func (a *Array) Get() interface{} {
	if a.size == 0 {
		return nil
	}
	return a.collection[a.size-1]
}

// Contains (Public) - Searches the Array O(n) for the first item that is contained in the Array
func (a *Array) Contains(item interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.collection[i] == item {
			return true
		}
	}
	return false
}

// Size (Public) - Return the size of the Array
func (a *Array) Size() int {
	return a.size
}

// String (Public) - returns the string version of the array if a funcion requests a string verison.
func (a *Array) String() string {
	result := "[ "
	for i := 0; i < a.size; i++ {
		result += fmt.Sprintf("%v ", a.collection[i])
	}
	return result + "]"
}

// RemoveItem (Public) - Removes a specific item. If it isn't found it exits.
// If found, moves all the other items left and nils last value for GC.
func (a *Array) RemoveItem(item interface{}) bool {
	if a.size == 0 {
		return false
	}
	removed := -1
	for i := 0; i < a.size; i++ {
		if a.collection[i] == item {
			removed = i
		}
	}
	if removed < 0 {
		return false
	}
	for i := removed; i < a.size-1; i++ {
		a.collection[i] = a.collection[i+1]
	}
	a.size--
	a.collection[a.size] = nil
	return true // To be implemented
}

// RemoveRandom (Public) - Gets a pseudo-random number using the time and removes that item.
// Then it goes through and makes sure that the rest of the items are moved left and there
// are no nil gaps
func (a *Array) RemoveRandom() interface{} {
	// if size is nothing you cannot remove anything
	if a.size == 0 {
		return nil
	}
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(a.size)
	choice := a.collection[random]
	for i := random; i < a.size-1; i++ {
		a.collection[i] = a.collection[i+1]
	}
	a.size--
	a.collection[a.size] = nil
	return choice
}

// GetRandom (Public) - Gets a pseudo-random number using the time and returns that item in the collection
func (a *Array) GetRandom() interface{} {
	if a.size == 0 {
		return nil
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return a.collection[rand.Intn(a.size)]
}

// ensureSpace (Private) - if the size and capacity of the array are the same, then a new array is created
// with twice the size of the array and the items contained are copied into the new array.
func (a *Array) ensureSpace() {
	if a.size == cap(a.collection) {
		new := new(Array).Init(cap(a.collection) * 2)
		new.size = a.size
		for i := 0; i < a.size; i++ {
			new.collection[i] = a.collection[i]
		}
		*a = *new
		new = nil
	}
}

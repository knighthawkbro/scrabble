package main

import (
	"fmt"
	"scrabble/array"
	"scrabble/list"
)

// Main interface that describes the encapsulated bag type
// Note that interface could mean anything, it is odd i know.
type bag interface {
	Add(i interface{}) error
	Remove() interface{}
	Get() interface{}
	Contains(i interface{}) bool
	Size() int
	RemoveItem(i interface{}) bool
	GetRandom() interface{}
	RemoveRandom() interface{}
}

// map of keys and values, values are a slice of strings, could be anything
var scrabbleLetters = map[int][]string{
	1:  []string{"K", "J", "X", "Q", "Z"},
	2:  []string{"B", "C", "M", "P", "F", "H", "V", "W", "Y"},
	3:  []string{"G"},
	4:  []string{"L", "S", "U", "D"},
	6:  []string{"N", "R", "T"},
	8:  []string{"O"},
	9:  []string{"A", "I"},
	12: []string{"E"},
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning scrable function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	lst := list.New()   // Creates a new linked-list called bag
	plist := list.New() // creates a new linked-list called player
	scrabble(lst, plist)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning scrable function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	arr := array.New()
	parray := array.New()
	scrabble(arr, parray)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as a list...")
	fmt.Println("\n*************************************************")
	list := list.New()
	driver(list)
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	array := array.New()
	driver(array)
}

func scrabble(b, p bag) {
	addLetters(b, scrabbleLetters)
	fmt.Printf("Bag letters size: %v\n", b.Size()) // fills the bag with the scrabble letters below
	selectLetters(b, p)                            // takes letters from bag and puts it into player
	fmt.Println("Your letters are:", p)            // prints the player list
	fmt.Printf("Bag letters size: %v\n", b.Size()) // prints the player list
}

// max letters per player
var lettersPerTurn = 7

// enumerates over the letters map to grab the key which becomes the count of each value
// then it goes over each value one by one times by the count for each letter
// adds those letters to which ever bag was passed in.
func addLetters(b bag, letters map[int][]string) {
	for count, value := range letters {
		for _, letter := range value {
			for i := 0; i < count; i++ {
				b.Add(letter)
			}
		}
	}
}

// runs a max of 7 times per the limit on player letters in the game
// removes the letter from the letters bag and places it in the
// players bag
func selectLetters(letters, player bag) {
	for i := 0; i < lettersPerTurn; i++ {
		letter := letters.RemoveRandom()
		player.Add(letter)
	}
}

func driver(words bag) {
	fruits := []string{"orange", "grape", "kiwi", "coconut", "lime"}
	for _, fruit := range fruits {
		words.Add(fruit)
	}
	fmt.Println("\nHere's what's in our bag:", words, words.Size())

	// seeing if bag contains item
	fmt.Printf("\nDoes our bag contain the word 'kiwi'? ")
	if words.Contains("kiwi") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Printf("Does our bag contain the word 'mango'? ")
	if words.Contains("mango") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// retrieving item
	fmt.Println("\nSelecting an item (always same)")
	fmt.Println(words.Get())
	fmt.Println(words.Get())

	// retrieving random item
	fmt.Println("\nSelecting a random item (varies)")
	fmt.Println(words.GetRandom())
	fmt.Println(words.GetRandom())

	// removing specific item
	words.RemoveItem("grape")
	fmt.Println("\nRemoving 'grape' from the bag\n", words)

	// removing item
	fmt.Println("\nRemoving an item (always an end one)")
	fmt.Println(words.Remove())
	fmt.Println(words)

	// remove random item
	fmt.Println("\nRemoving a random item")
	fmt.Println(words.RemoveRandom())
	fmt.Println(words)

	// testing methods on empty bag
	fmt.Println("Let's empty the bag")
	words.Remove()
	words.Remove()
	fmt.Println("Trying to get a random item (should be nil)\n", words.GetRandom())
	fmt.Println("Trying to remove a random item (should be nil)\n", words.RemoveRandom())
	fmt.Println("Trying to remove 'kiwi' (should be false)\n", words.RemoveItem("kiwi"))
}

package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favCards := []int{2, 6, 9}
	return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	}
	for i := range slice {
		if i == index {
			return slice[i]
		}
	}
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	for i := range slice {
		if i == index {
			slice[i] = value
		}
	}
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
	}
	return slice
	panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var slice2 []int
	for _, value := range values {
		slice2 = append(slice2, value)
	}
	slice2 = append(slice2, slice...)
	return slice2
	panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var slice1, slice2 []int
	if index >= len(slice) || index < 0 {
		return slice
	}
	if index == 0 {
		slice = slice[1:]
		return slice
	} else if index == len(slice)-1 {
		slice = slice[:len(slice)-1]
		return slice
	} else {
		slice1 = slice[0:index]
		slice2 = slice[index+1:]
		slice = append(slice1, slice2...)
		return slice
	}
	panic("Please implement the RemoveItem function")
}

func main() {

	// card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4
	// fmt.Println(card)

	// slice := []int{3, 2, 6, 4, 8}
	// cards := PrependItems(slice)
	// fmt.Println(cards)

	cards := RemoveItem([]int{0, 1, 2, 3, 4, 5}, 2)
	fmt.Println(cards)
	// Output: [3 2 4 8]

}

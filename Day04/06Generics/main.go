package main

import "fmt"

func printsomething[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
}

func maplkeys[K comparable, v any](m map[K]v) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	// s := []int{1, 2, 3, 4, 5}
	// printsomething(s)
	// fmt.Println()
	// p := []string{"omkar", "patil", "you", "can", "do", "it"}
	// printsomething(p)
	// fmt.Println()

	// m := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	// fmt.Println(maplkeys(m))

}

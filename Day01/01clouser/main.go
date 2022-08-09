package main

import "fmt"

func newcounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	counter := newcounter()
	fmt.Println(counter())
	fmt.Println(counter())
}

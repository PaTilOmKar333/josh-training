package main

import "fmt"

func byvalue(n int) {
	n = 5
}

func byrefrance(n *int) {
	*n = 5
}

func main() {
	n := 1
	fmt.Println("n:", n)
	byvalue(n)
	fmt.Println("n(byvalue):", n)
	byrefrance(&n)
	fmt.Println("n(byrefrance)", n)
	fmt.Println(&n)

}

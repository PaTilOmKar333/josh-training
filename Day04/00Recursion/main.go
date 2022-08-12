package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// var num int
	// fmt.Printf("enter number:")
	// fmt.Scan(&num)
	// f := factorial(num)
	// fmt.Println(f)

	var fib func(n int) int

	fib = func(n int) int {
		fmt.Println(n)
		if n < 2 {
			return n
		}
		fmt.Printf("fib(n-1):%d", fib(n-1))
		fmt.Printf("fib(n-2):%d", fib(n-2))
		return fib(n-1) + fib(n-2)

	}
	fmt.Println(fib(7))

}

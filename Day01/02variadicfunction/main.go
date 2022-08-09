package main

import "fmt"

func sum(nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)

}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)
	sum(1, 2, 3, 4, 5)
}

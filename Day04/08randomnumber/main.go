package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixMicro())
	randomNumber := rand.Intn(100)
	//rand.Seed(time.Now().UnixNano())

	fmt.Println(randomNumber)
	var num int

	for {
		fmt.Println("Enter number:")
		fmt.Scan(&num)
		if randomNumber == num {
			fmt.Println("randomNumber == num")
			break
		} else if num < 0 {
			break
		} else if randomNumber > num {
			fmt.Println("randomNumber > num")
		} else {
			fmt.Println("randomNumber < num")
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func func1(ch chan string) {
	time.Sleep(1 * time.Second)
	//ch <- "func1 excuted successfully"
}
func func2(ch chan string) {
	time.Sleep(2 * time.Second)
	//ch <- "func2 excuted successfully"
}
func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func1(ch1)
	go func2(ch2)

	select {
	case res01 := <-ch1:
		fmt.Println(res01)
	case res02 := <-ch2:
		fmt.Println(res02)
		// default:
		// 	fmt.Println("Not Found")
	}

}

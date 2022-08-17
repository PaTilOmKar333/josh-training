package main

import "fmt"

func main() {

	message := make(chan string, 2)

	message <- "omkar"
	message <- "patil"

	fmt.Println(<-message)
	fmt.Println(<-message)

}

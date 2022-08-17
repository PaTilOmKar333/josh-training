package main

import "fmt"

func ping(pings chan<- string, str string) {
	pings <- str
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func pongwe(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hi, omkar here")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

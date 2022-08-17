package main

import (
	"fmt"
	"time"
)

func iscomplate(complate chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	complate <- true
}

func main() {
	complate := make(chan bool, 1)
	go iscomplate(complate)
	<-complate
}

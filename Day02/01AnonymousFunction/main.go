package main

import "fmt"

func something() func() {
	n := 1

	return func() {
		fmt.Println(n)
		n++

	}
}

func anotherfuncrtion(name string) {
	name = "patil"
}

func main() {
	// a := func() {
	// 	fmt.Println("hi Josh")
	// }
	// fmt.Printf("type of a:%T", a)
	//h := something()
	// h()
	// h()
	// h()
	// something()()
	// something()()
	// something()()
	//name := "omkar"
	//b := callB(2)
	name := "omkar"
	anotherfuncrtion(name)
	fmt.Println(name)

	//fmt.Println(call("omkar"))
}

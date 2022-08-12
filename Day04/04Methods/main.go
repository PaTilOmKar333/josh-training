package main

import "fmt"

type rectangle struct {
	height, width int
}

func (r rectangle) area() int {
	return r.height * r.width
}

func (r rectangle) perimeter() int {
	return 2 * (r.height + r.width)
}

func main() {
	r := rectangle{height: 10, width: 20}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perimeter())
}

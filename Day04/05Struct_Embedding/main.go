package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num: %v", b.num)
}

type container struct {
	base
	name string
}

func main() {

	co := container{
		base: base{
			num: 10,
		},
		name: "omkar",
	}
	b := base{
		num: 20,
	}

	fmt.Println(co, co.base, co.name, co.num, co.base.num)
	fmt.Println(co.describe())
	fmt.Println(co.base.describe())
	fmt.Println(b.num)
	fmt.Println(b.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println(d.describe())

}

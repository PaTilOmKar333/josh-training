package main

import "fmt"

type employee struct {
	eid    int
	name   string
	salary float32
}

func newEmployee(name string) *employee {
	E := employee{name: name}
	E.eid = 10
	E.salary = 10000.00
	return &E
}

func main() {

	var e1 employee
	fmt.Println(e1)
	e1.eid = 1
	e1.name = "omkar"
	e1.salary = 9999999.99
	fmt.Println(e1)

	fmt.Println(employee{2, "patil", 99999.8989})
	fmt.Println(employee{eid: 3, name: "adhiraj", salary: 12344567.8765})
	fmt.Println(employee{eid: 4, name: "dhanraj"})
	fmt.Println(employee{eid: 4})

	EE := newEmployee("aditi")
	fmt.Println(*EE)

	Ep := employee{eid: 3, name: "Vishvraj", salary: 12344567.8765}
	ptrEp := &Ep
	fmt.Println(ptrEp.name)
	ptrEp.salary = 999999.1234

	fmt.Println(ptrEp)
	fmt.Println(&ptrEp)
	fmt.Println(*ptrEp)

}

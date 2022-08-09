package main

import (
	"fmt"
)

const TechDepartMent = "Tech"
const DigiDepartMent = "Digi"

type department struct {
	dep string
}

type employee struct {
	name       string
	age        int
	salary     float64
	department map[string]department
}

func Newemployee(Name string, Age int, Salary float64) employee {

	emp := employee{
		name:   Name,
		age:    Age,
		salary: Salary,
	}
	emp.department = make(map[string]department)
	return emp

}

func Newdepartment(departmnet string) department {

	return department{
		dep: departmnet,
	}
}

func GetdepartmentName(TechDepartMent string, e employee) (departmentName department, err error) {
	departmentName, ok := e.department[TechDepartMent]
	if !ok {
		err = fmt.Errorf("doesnot exist department")
	}
	return
}

func Printdepartment(e employee) {
	for i, result := range e.department {
		fmt.Println(i, ":", result)
	}
}

func DeleteDepartment(tye string, e employee) (err error) {
	_, ok := e.department[tye]
	if !ok {
		err = fmt.Errorf("doesnot exist department")
	}
	delete(e.department, tye)
	return
}

func main() {

	e1 := Newemployee("omkar", 27, 90000)
	e1.department[TechDepartMent] = Newdepartment("Golang")

	departmentName, err := GetdepartmentName(TechDepartMent, e1)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(e1.name, e1.age, e1.salary, departmentName.dep)
	e1.department[DigiDepartMent] = Newdepartment("Devops")
	fmt.Println(e1.name, e1.age, e1.salary, departmentName.dep)

	Printdepartment(e1)

	DeleteDepartment(TechDepartMent, e1)

	Printdepartment(e1)

}

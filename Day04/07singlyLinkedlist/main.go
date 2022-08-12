package main

import "fmt"

type linkendlist struct {
	number int
	next   *linkendlist
}

func (l *linkendlist) insertinto(num int) {
	temp := &linkendlist{}
	temp.number = num
	temp.next = nil
	if l == nil {
		l = temp
	} else {
		p := &linkendlist{}
		p = l
		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}
}
func (l *linkendlist) Dispalylinkedlist() {
	p := &linkendlist{}
	p = l.next
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
	fmt.Println()
}

func main() {
	head := new(linkendlist)
	head.insertinto(10)
	head.insertinto(20)
	head.insertinto(30)
	head.insertinto(40)
	head.insertinto(50)
	head.Dispalylinkedlist()
}

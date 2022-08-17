package main

import (
	"errors"
	"fmt"
)

func f1(n int) (int, error) {
	if n == 42 {
		return n, errors.New("its error")
	} else {
		return n, nil
	}
}

type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.msg)
}

func f2(n int) (int, error) {
	if n == 42 {
		return -1, &argError{n, "can't work with it"}
	} else {
		return n, nil
	}
}

func main() {

	for _, n := range []int{7, 23, 42, 65} {
		if r, e := f1(n); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}

	}

	fmt.Println("*************************************")

	for _, n := range []int{7, 23, 42, 65} {
		if r, e := f2(n); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}

	}
	fmt.Println("*************************************8")

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	}
}

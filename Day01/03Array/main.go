package main

import "fmt"

func main() {
	var array1 = [4]int{1, 2, 3, 4}
	var array2 = [4]int{1, 2, 3, 4}
	array3 := [4]string{"a", "b", "c", "d"}

	fmt.Println(array1, array2, array3)
	b := array1 == array2
	fmt.Println(b)
	if array1 == array2 {
		fmt.Println("comparable")
	}

	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}
}

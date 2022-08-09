package main

import "fmt"

func comparetwoslices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {

	// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice01 := arr[0:5]
	// slice02 := arr[2:7]
	// slice03 := arr[3:9]
	// slice04 := arr[7:]
	// slice05 := arr[:7]
	// fmt.Println("arr[0:5]", arr, len(slice01), cap(slice01), slice01) //1,5 //for first value index should be start from 0
	// fmt.Println("arr[2:7]", arr, len(slice02), cap(slice02), slice02) //3,7 //for second value index should be start from 1
	// fmt.Println("arr[3:9]", arr, len(slice03), cap(slice03), slice03) //4,9
	// fmt.Println("arr[7: ]", arr, len(slice04), cap(slice04), slice04) //8,10
	// fmt.Println("arr[:7 ]", arr, len(slice05), cap(slice05), slice05) //1,7

	// // slice03[4] = 77
	// // fmt.Println("arr[0:5]", arr, slice01) //1,5 //for first value index should be start from 0
	// // fmt.Println("arr[2:7]", arr, slice02) //3,7 //for second value index should be start from 1
	// // fmt.Println("arr[3:9]", arr, slice03) //4,9
	// // fmt.Println("arr[7: ]", arr, slice04) //8,10
	// // fmt.Println("arr[:7 ]", arr, slice05) //1,7

	// slice06 := slice01
	// fmt.Println("arr[0:5]", arr, len(slice06), cap(slice06), slice06) //1,5 //for first value index should be start from 0

	// //fmt.Println(slice01 == slice06)
	// fmt.Println(comparetwoslices(slice01, slice02))

	// slice07 := make([]int, 5, 6)
	// fmt.Println("slice07", arr, len(slice07), cap(slice07), slice07)

	// sl1 := []int{1, 2, 3, 4, 5}
	// sl2 := sl1[1:4]

	// fmt.Println(sl1)
	// fmt.Println(sl2)

	// //sl2 = append(sl2, 10, 11, 12, 13, 14, 15)
	// sl2 = append(sl2, 10)

	// sl3 := append(sl1, sl2...)

	// fmt.Println(sl1)
	// fmt.Println(sl2)
	// fmt.Println(sl3)

	a := make([]int, 5)
	fmt.Println(len(a), cap(a))

	b := make([]int, 0, 5)
	fmt.Println(len(b), cap(b))

	c := b[:2]
	fmt.Println(len(c), cap(c))

	d := c[2:5]
	fmt.Println(len(d), cap(d))
}

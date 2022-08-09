package calculator

import "fmt"

func Sum(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

func Pricecalculation(price, num int) int {
	totalprice := price * num
	return totalprice
}

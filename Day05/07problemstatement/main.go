// abc ,a,ab,bc module 3

package main

import (
	"fmt"
	"sort"
)

func sortString(strArry []string) {
	newStrArray := sort.StringSlice(strArry)
	//newStrArray := sort.Strings(strArry)
	fmt.Println("newStrArray:", newStrArray)

	for i := 0; i < len(strArry); i++ {
		//	fmt.Println(strArry[i], ":", len(strArry[i]), ":", len(strArry[i])%3)

		//		b := len(strArry[i])%3 == 0
		//	fmt.Println("b:", b)
		if len(strArry[i])%3 == 0 {
			fmt.Println("0:", strArry[i])
			fmt.Println(newStrArray)
			newStrArray = append(newStrArray, strArry[i])
			fmt.Println(newStrArray)
		} else if len(strArry[i])%3 == 1 {
			fmt.Println("1:", strArry[i])
			fmt.Println(newStrArray)
			//	newStrArray1 = append(newStrArray1, strArry[i])
		} else {
			//	newStrArray2 = append(newStrArray2, strArry[i])
		}
	}
	fmt.Println("newStrArray", newStrArray)
	//fmt.Println("newStrArray1", newStrArray1)
	//fmt.Println("newStrArray2", newStrArray2)

}
func ok() {

}

func main() {
	str := []string{"omkar", "patil", "Hows", "the", "josh"}
	sortString(str)
}

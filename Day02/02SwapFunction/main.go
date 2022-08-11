package main

import (
	"fmt"
	"strings"
)

// func SwapByValue(a, b int) (int, int) {
func SwapByValue(a, b int) {
	tmp := a
	a = b
	b = tmp

	//return a, b
}

func SwapByRefrance(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	//return *a, *b
}

func nakedreturn() (firstname, lastname string) {
	firstname = "omkar"
	lastname = "patil"
	return
}
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	//panic("Please implement the CanFreePrisoner() function")
	fmt.Printf("knightIsAwake:%v, archerIsAwake:%v, prisonerIsAwake:%v, petDogIsPresent:%v\n", knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent)
	fmt.Printf("knightIsAwake:%v, archerIsAwake:%v, prisonerIsAwake:%v, petDogIsPresent:%v\n", !knightIsAwake, !archerIsAwake, !prisonerIsAwake, !petDogIsPresent)

	if petDogIsPresent {
		if !archerIsAwake {
			return true
		} else {
			return false
		}

	} else {
		if (knightIsAwake || archerIsAwake) && prisonerIsAwake {
			return true
		} else {
			return false
		}
	}
}

func WelcomeMessage(customer string) string {
	//panic("Please implement the WelcomeMessage() function")
	capitalName := strings.ToUpper(customer)
	return "Welcome to the Tech Palace," + capitalName
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	//panic("Please implement the AddBorder() function")
	starstring := ""
	i := 1
	for i <= numStarsPerLine {
		starstring += "*"
		i++
	}
	return starstring + "\n" + welcomeMsg + "\n" + starstring
}

func CleanupMessage(oldMsg string) string {
	//panic("Please implement the CleanupMessage() function")
	newMsg := ""
	for _, v := range oldMsg {
		//fmt.Println(string(v))
		if string(v) == "*" || string(v) == "\n" || string(v) == "\t" {

		} else {
			newMsg += string(v)
			//fmt.Print(newMsg)
		}
	}
	return strings.Trim(newMsg, "  ")

	//return newMsg
}

func main() {
	// a, b:=SwapByValue(10, 20)
	// fmt.Println("SwapByValue:a,b", a, b)
	// a1, b1 := 10, 20

	// a1, b1 = SwapByRefrance(&a1, &b1)
	// fmt.Println("SwapByRefrance:a,b", a1, b1)
	//--------------------------
	// var p int = 10
	// var q int = 20
	// fmt.Printf("p = %d and q = %d", p, q)

	// // call by values
	// SwapByRefrance(&p, &q)
	// fmt.Printf("\np = %d and q = %d", p, q)
	//---------------------------

	// a:= true
	// b:= false
	//fmt.Println(nakedreturn())

	// knightIsAwake := false
	// archerIsAwake := false
	// prisonerIsAwake := false
	// dogIsPresent := false

	// fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, dogIsPresent))

	//fmt.Println(WelcomeMessage("omkar"))

	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`

	//fmt.Println(AddBorder("khedkar ", 1))
	fmt.Println(CleanupMessage(message))

}

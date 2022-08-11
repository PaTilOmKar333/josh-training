package main

import (
	"fmt"
	"math"
	"strconv"
)

func Welcome(name string) string {
	//panic("Please implement the Welcome function")
	return "Welcome to my party, " + name + "!"
}

func HappyBirthday(name string, age int) string {
	//panic("Please implement the HappyBirthday function")
	//strings(age)
	Age := strconv.Itoa(age)
	return "Happy birthday " + name + "! You are now " + Age + " years old!"
}
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {

	Table := strconv.Itoa(table)
	rem := 3 - len(Table)
	var zerostring string
	i := 0
	for i < rem {
		zerostring += "0"
		i++
	}
	NewTable := zerostring + Table

	distance = math.Round(distance*100) / 100
	fmt.Println(distance)
	Distance := strconv.FormatFloat(distance, 'f', 1, 32)
	return "Welcome to my party, " + name + "! \nYou have been assigned to table " + NewTable + ". Your table " + direction + ", exactly " + Distance + " meters from here.\nYou will be sitting next to " + neighbor + "."
}

func main() {
	//fmt.Println(HappyBirthday("omkar", 27))

	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 9.2394381))
	//AssignTable("Christiane", 2, "Frank", "on the left", 23.7834298)
	// =>
	// Welcome to my party, Christiane!
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	// You will be sitting next to Frank.

	//"Welcome to my party, Chihiro!\nYou have been assigned to table 022. Your table is straight ahead, exactly 9.2 meters from here.\nYou will be sitting next to Akachi Chikondi."
	//"Welcome to my party, Chihiro!\nYou have been assigned to table 022. Your table isstraight ahead, exactly 9.2 meters from here.\nYou will be sitting next to Akachi Chikondi."
}

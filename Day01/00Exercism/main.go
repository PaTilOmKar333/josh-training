package main

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	RemainingOvenTime := OvenTime - actual
	return RemainingOvenTime
}

func PreparationTime(numberOfLayers int) int {
	PreparationTime := numberOfLayers * 2
	return PreparationTime
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	PreparationTime := numberOfLayers * 2
	RemainingOvenTime := OvenTime - actualMinutesInOven
	ElapsedTime := PreparationTime + RemainingOvenTime

	return ElapsedTime
}

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	CalculateWorkingCarsPerHour := float64(productionRate) * successRate / 100
	// panic("CalculateWorkingCarsPerHour not implemented")
	return CalculateWorkingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	CalculateWorkingCarsPerMinute := float64(productionRate) * successRate / (100 * 60)
	//panic("CalculateWorkingCarsPerMinute not implemented")
	return int(CalculateWorkingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	//panic("CalculateCost not implemented")
	indivudalCarCount := carsCount % 10
	carsCount -= indivudalCarCount
	groupOftenCount := carsCount / 10

	CalculateCost := (groupOftenCount * 95000) + (indivudalCarCount * 10000)
	return uint(CalculateCost)
}

func main() {
	//fmt.Println(RemainingOvenTime(10))
	//fmt.Println(PreparationTime(10))
	//fmt.Println(ElapsedTime(2, 10))
	//fmt.Println(CalculateWorkingCarsPerHour(0, 100))
	a := 37
	ekak := a % 10
	dashak := (a - ekak) / 10
	fmt.Println(dashak)

}

package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(float64(productionRate) * successRate * 0.01)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const costPerGroup = 95000
    const costPerCar = 10000
    groups := int(carsCount / 10)
    remainingCars := carsCount - (groups * 10)
    groupedCarsCost := groups * costPerGroup
    remainingCarsCost := remainingCars * costPerCar
	return uint(groupedCarsCost + remainingCarsCost)
}

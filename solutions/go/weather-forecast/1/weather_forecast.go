// Package weather forecasts weather of a supplied city in conditions described.
package weather

// CurrentCondition stores current condition of the city.
var CurrentCondition string
// CurrentLocation stores current city.
var CurrentLocation string

// Forecast takes city and condition and returns a report describing it.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

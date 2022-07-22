// Package weather contains ForestCast.
package weather

// CurrentCondition holds the current weather.
var CurrentCondition string

// CurrentLocation holds the location of the weather.
var CurrentLocation string

// Forecast returns the weather for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

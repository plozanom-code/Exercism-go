// Package weather provides tools to ascertain the weather in a particular location.
package weather

var (
    // CurrentCondition represents the weather in a set current location.
	CurrentCondition string
	// CurrentLocation represents the place we want to know the weather for.
    CurrentLocation  string
)

// Forecast returns a string that says the weather of the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

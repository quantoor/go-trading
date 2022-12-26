package util

import (
	"fmt"
	"math"
	"strings"
)

func RoundWithPrecision(x float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Floor(x*output) / output
}

func TickToPrecision(tick float64) int {
	decimalPlaces := fmt.Sprintf("%.16f", tick)
	decimalPlaces = strings.Replace(decimalPlaces, "0.", "", -1) // remove 0.
	decimalPlaces = strings.TrimRight(decimalPlaces, "0")        // remove trailing 0s
	return len(decimalPlaces)
}

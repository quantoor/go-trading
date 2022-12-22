package util

import "math"

func RoundWithPrecision(x float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Floor(x*output) / output
}

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 2.0
	itr := 0

	for diff := 1.0; math.Abs(diff) > 1e-10; itr++ {
		diff = (z*z - x) / (2 * z)
		z -= diff
	}
	fmt.Println("Iterations : ", itr)

	return z
}

func main() {
	fmt.Println("Using Newton's method : ", Sqrt(2))
	fmt.Println("Using math.Sqrt       : ", math.Sqrt(2))
}

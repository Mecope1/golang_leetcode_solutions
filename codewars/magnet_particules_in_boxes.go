package kata

import "math"

func Doubles(maxk, maxn int) float64 {
	// your code

	summation := 0.0

	for k := 1; k <= maxk; k++ {
		for n := 1; n <= maxn; n++ {
			summation += 1.0 / (float64(k) * math.Pow(float64(n+1), float64(2*k)))
		}
	}
	return summation
}
package chap2_3

import (
	"math"
)

func P2351(x int) int {
	return int(math.Log10(float64(x)))
}

func P2352(n int) int {
	return int(math.Log2(float64(16*n)) - math.Log2(float64(n)))
}

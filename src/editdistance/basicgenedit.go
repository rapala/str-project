package editdistance

import (
	"math"
)

func MakeBasicGenEdit(G [][]string, c []float64) (func(string, string) float64) {
	cost := func(A string, B string, pi int, d [][]float64) float64 {
		a := len(A) - len(G[pi][0])
		b := len(B) - len(G[pi][1])
		if a >= 0 && b >= 0 && A[a:] == G[pi][0] && B[b:] == G[pi][1] {
			return d[b][a]+c[pi]
		}
		return math.Inf(1)
	}

	minCost := func(A string, B string, d [][]float64) float64 {
		min := math.Inf(1)
		if len(A) > 0 && len(B) > 0 && A[len(A)-1] == B[len(B)-1] {
			min = d[len(B)-1][len(A)-1]
		}
		for pi, _ := range G {
			min = math.Fmin(min, cost(A, B, pi, d))
		}
		return min
	}

	return func(A string, B string) float64 {
		d := makeMatrix(len(B)+1, len(A)+1)

		for x := 0; x <= len(A); x++ {
			for y := 0; y <= len(B); y++ {
				if x == 0 && y == 0 {
					d[y][x] = 0
				} else {
					d[y][x] = minCost(A[:x], B[:y], d)
				}
			}
		}

		return d[len(B)][len(A)]
	}
}
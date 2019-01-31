package math

import "math"

// Problem: Count the number of perfect squares between two numbers

func countPerfectSquares(x, y int64) int {
  result := math.Floor(math.Sqrt(float64(y))) - math.Ceil(math.Sqrt(float64(x))) + 1
  return int(result)
}

package math

import "math"

// Problem: Count the number of perfect squares between two numbers

func countPerfectSquares(x, y int64) int {
  result := math.Floor(math.Sqrt(float64(y))) - math.Ceil(math.Sqrt(float64(x))) + 1
  return int(result)
}

// Problem: Find perfect squares between two numbers
// assume the numbers will always be positive

// solve using (n + 1)^2 = n^2 + 2n + 1
// where first n = int(sqrt(x)) and iterate until n < k
func findPerfectSquares(x, y int64) []int64 {
  n := int64(math.Sqrt(float64(x)))
  k := int64(math.Sqrt(float64(y)))
  var squares []int64
  
  // exclude y
  if y-k*k == 0 {
    k--
  }
  
  for square := n * n; n < k; n++ {
    square += 2*n + 1
    squares = append(squares, square)
  }
  return squares
}

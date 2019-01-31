package math

import (
  "testing"
)

func TestCountSquares(t *testing.T) {
  expected := 2
  actual := countPerfectSquares(2, 10)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestCountSquaresSameNumbers(t *testing.T) {
  expected := 0
  actual := countPerfectSquares(2, 2)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestCountSquaresSquareNumbers(t *testing.T) {
  expected := 0
  actual := countPerfectSquares(4, 9)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

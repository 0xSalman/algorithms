package math

import (
  "reflect"
  "testing"
)

func TestCountPerfectSquares(t *testing.T) {
  expected := 2
  actual := countPerfectSquares(2, 10)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestCountPerfectSquares_SameNumbers(t *testing.T) {
  expected := 0
  actual := countPerfectSquares(2, 2)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestCountPerfectSquares_SquaredNumbers(t *testing.T) {
  expected := 0
  actual := countPerfectSquares(4, 9)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestFindPerfectSquares(t *testing.T) {
  expected := []int64{4, 9, 16, 25}
  actual := findPerfectSquares(3, 30)
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestFindPerfectSquares_SameNumbers(t *testing.T) {
  var expected []int64
  actual := findPerfectSquares(3, 3)
  if actual != nil || len(actual) > 0 {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestFindPerfectSquares_SquaredNumbers(t *testing.T) {
  expected := []int64{9, 16}
  actual := findPerfectSquares(4, 25)
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

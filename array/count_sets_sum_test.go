package array

import "testing"

func TestCountSetsSimple(t *testing.T) {
  numbers := []int{4, 2, 10, 6}
  expected := 2
  actual := countSetsEqualingSum(numbers, 16)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestCountSetsWithDuplicates(t *testing.T) {
  numbers := []int{4, 2, 4, 10, 10, 6}
  expected := 3
  actual := countSetsEqualingSum(numbers, 20)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

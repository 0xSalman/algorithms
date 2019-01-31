package array

import "testing"

func TestSmallLarge(t *testing.T) {
  numbers := []int{1, 4, 32, 99}
  expectedSmall := 1
  expectedLarge := 99
  small, large := findSmallAndLarge(numbers)
  if small != expectedSmall && large != expectedLarge {
    t.Errorf("Got small: %d, wanted: %d\n", small, expectedSmall)
    t.Errorf("Got large: %d, wanted: %d\n", large, expectedLarge)
  }
}

func TestSmallLargeWithNegative(t *testing.T) {
  numbers := []int{-100, -50, -501, 101, 403, 1001}
  expectedSmall := -501
  expectedLarge := 1001
  small, large := findSmallAndLarge(numbers)
  if small != expectedSmall && large != expectedLarge {
    t.Errorf("Got small: %d, wanted: %d\n", small, expectedSmall)
    t.Errorf("Got large: %d, wanted: %d\n", large, expectedLarge)
  }
}

func TestSmallLargeSizeOne(t *testing.T) {
  numbers := []int{-3}
  expectedSmall := -3
  expectedLarge := -3
  small, large := findSmallAndLarge(numbers)
  if small != expectedSmall && large != expectedLarge {
    t.Errorf("Got small: %d, wanted: %d\n", small, expectedSmall)
    t.Errorf("Got large: %d, wanted: %d\n", large, expectedLarge)
  }
}

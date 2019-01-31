package array

import (
  "testing"
)

func TestLargestPairSum(t *testing.T) {
  numbers := []int{1, 15, 10, 6, 9}
  expected := 25
  actual := findLargestPairSum(numbers)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestLargestPairSumDuplicates(t *testing.T) {
  numbers := []int{1, 15, 10, 6, 9, 15}
  expected := 30
  actual := findLargestPairSum(numbers)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestLargestPairSumMultipleDuplicates(t *testing.T) {
  numbers := []int{20, 15, 20, 30, 30, 10, 6, 30, 9, 15}
  expected := 50
  actual := findLargestPairSum(numbers)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestLargestPairSumSameNumber(t *testing.T) {
  numbers := []int{1, 1, 1, 1, 1}
  expected := 2
  actual := findLargestPairSum(numbers)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestLargestPairEmptyList(t *testing.T) {
  var numbers []int
  expected := 0
  actual := findLargestPairSum(numbers)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestLargestPairListSizeOne(t *testing.T) {
  numbers := []int{5}
  expected := 5
  actual := findLargestPairSum(numbers)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestLargestPairListSizeTwo(t *testing.T) {
  numbers := []int{5, 10}
  expected := 15
  actual := findLargestPairSum(numbers)
  if actual != expected {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

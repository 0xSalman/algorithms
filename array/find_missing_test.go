package array

import (
  "reflect"
  "testing"
)

func TestMissingOne(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
  expected := []int{6}
  actual := findMissingNumbers(numbers, 1, 10)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestMissingMultiple(t *testing.T) {
  numbers := []int{1, 2, 3, 10}
  expected := []int{4, 5, 6, 7, 8, 9}
  actual := findMissingNumbers(numbers, 1, 10)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestMissingMultipleWithNegatives(t *testing.T) {
  numbers := []int{-5, -4, -3, 1, 2, 3}
  expected := []int{-2, -1, 0, 4, 5}
  actual := findMissingNumbers(numbers, -5, 5)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestMissingWithDuplicates(t *testing.T) {
  numbers := []int{51, 53, 54, 54, 57}
  expected := []int{52, 55, 56}
  actual := findMissingNumbers(numbers, 51, 57)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestMissingNotSorted(t *testing.T) {
  numbers := []int{65, 67, 69, 68, 67, 62, 63, 65, 61, 70}
  expected := []int{60, 64, 66}
  actual := findMissingNumbers(numbers, 60, 70)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

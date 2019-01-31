package array

import (
  "reflect"
  "sort"
  "testing"
)

func TestDuplicateOne(t *testing.T) {
  numbers := []int{1, 2, 3, 3, 4, 6}
  expected := []int{3}
  actual := findDuplicates(numbers)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestDuplicateMultiple(t *testing.T) {
  numbers := []int{1, 2, 6, 21, 21, 17, 17, 4, 6}
  expected := []int{6, 17, 21}
  actual := findDuplicates(numbers)
  sort.Ints(actual)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

func TestDuplicateNone(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5, 6, 7}
  expected := []int{}
  actual := findDuplicates(numbers)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %d, wanted: %d\n", actual, expected)
  }
}

package array

import (
  "reflect"
  "testing"
)

func TestPairFaster(t *testing.T) {
  numbers := []int{0, 2, 4, 3, 6, 5, 7, 8, 9, 11, -15}
  expected := []pair{
    {5, 6},
    {7, 4},
    {8, 3},
    {9, 2},
    {11, 0},
  }
  actual := findPairsFaster(numbers, 11)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestPairFasterNone(t *testing.T) {
  numbers := []int{1, 2, 3, 9}
  expected := make([]pair, 0)
  actual := findPairsFaster(numbers, 8)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
  
}

func TestPairFasterEmpty(t *testing.T) {
  var numbers []int
  expected := make([]pair, 0)
  actual := findPairsFaster(numbers, 8)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestPairFasterWithDuplicates(t *testing.T) {
  numbers := []int{2, 4, 3, 5, 6, -2, 4, 7, 8, 9}
  expected := []pair{
    {3, 4},
    {5, 2},
    {9, -2},
  }
  actual := findPairsFaster(numbers, 7)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestPairEfficiently(t *testing.T) {
  numbers := []int{0, 2, 4, 3, 6, 5, 7, 8, 9, 11, -15}
  expected := []pair{
    {0, 11},
    {2, 9},
    {3, 8},
    {4, 7},
    {5, 6},
  }
  actual := findPairsEfficiently(numbers, 11)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestPairEfficientlyNone(t *testing.T) {
  numbers := []int{1, 2, 3, 9}
  expected := make([]pair, 0)
  actual := findPairsEfficiently(numbers, 8)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestPairEfficientlyEmpty(t *testing.T) {
  var numbers []int
  expected := make([]pair, 0)
  actual := findPairsEfficiently(numbers, 8)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

func TestPairEfficientlyWithDuplicates(t *testing.T) {
  numbers := []int{2, 4, 3, 5, 6, -2, 4, 7, 8, 9}
  expected := []pair{
    {-2, 9},
    {2, 5},
    {3, 4},
  }
  actual := findPairsEfficiently(numbers, 7)
  if reflect.DeepEqual(actual, expected) == false {
    t.Errorf("Got: %v, wanted: %v\n", actual, expected)
  }
}

package array

import "math"

// Problem: given an integer array, find the smallest
// and largest number
// assume array is not sorted and has at least one value

func findSmallAndLarge(numbers []int) (int, int) {
  if len(numbers) == 1 {
    return numbers[0], numbers[0]
  }
  
  smallest := numbers[0]
  largest := numbers[0]
  for _, num := range numbers {
    smallest = int(math.Min(float64(smallest), float64(num)))
    largest = int(math.Max(float64(largest), float64(num)))
  }
  
  return smallest, largest
}

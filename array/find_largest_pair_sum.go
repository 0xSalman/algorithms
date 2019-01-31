package array

import "math"

// Problem: given an integer array, find the largest pair sum
// assume the array is not sorted and does not contain duplicates
// For example, the largest pair sum in {1, 15, 10, 6, 9} is 25

// The simplest solution would be to sort the array
// and sum the last two numbers but complexity for
// this algorithm would be O(nlogn).
// This can be solved with O(n) complexity

func findLargestPairSum(numbers []int) int {
  if len(numbers) == 0 {
    return 0
  }
  
  if len(numbers) == 1 {
    return numbers[0]
  }
  
  if len(numbers) == 2 {
    return numbers[0] + numbers[1]
  }
  
  first := int(math.Max(float64(numbers[0]), float64(numbers[1])))
  second := int(math.Min(float64(numbers[0]), float64(numbers[1])))
  
  for i := 2; i < len(numbers); i++ {
    if numbers[i] > first {
      second = first
      first = numbers[i]
    } else if numbers[i] != first && numbers[i] > second {
      second = numbers[i]
    }
  }
  
  return first + second
}

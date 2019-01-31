package array

import (
  "fmt"
  "math"
)

// Problem: given an integer array, find duplicates
// assume array is not sorted

func findDuplicates(numbers []int) []int {
  duplicates := make([]int, 0)
  if len(numbers) <= 1 {
    return duplicates
  }
  
  numbersMap := make(map[int]struct{})
  for _, num := range numbers {
    if _, ok := numbersMap[num]; ok {
      duplicates = append(duplicates, num)
    } else {
      numbersMap[num] = struct{}{}
    }
  }
  
  return duplicates
}

func miniMaxSum(arr []int32) {
  var numbers [5]int64
  for i, value := range arr {
    numbers[i] = int64(value)
  }
  
  sum := numbers[1] + numbers[2] + numbers[3] + numbers[4]
  min, max := minMax(0, 0, sum)
  sum = numbers[0] + numbers[2] + numbers[3] + numbers[4]
  min, max = minMax(min, max, sum)
  sum = numbers[0] + numbers[1] + numbers[3] + numbers[4]
  min, max = minMax(min, max, sum)
  sum = numbers[0] + numbers[1] + numbers[2] + numbers[4]
  min, max = minMax(min, max, sum)
  sum = numbers[0] + numbers[1] + numbers[2] + numbers[3]
  min, max = minMax(min, max, sum)
  fmt.Printf("%d %d\n", min, max)
}

func minMax(min, max int64, sum int64) (int64, int64) {
  if min == 0 {
    return sum, sum
  } else {
    return int64(math.Min(float64(sum), float64(min))),
      int64(math.Max(float64(sum), float64(max)))
  }
}

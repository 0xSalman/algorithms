package array

import "strconv"

// Problem: given an integer array N, count all sets of numbers
// whose sum equals M where M >= 1
// assume array is not sorted and may contain duplicates
// For example, given N=[2, 4, 6, 10] and M=16 the sets whose
// sum equals 16 are {2, 4, 10} and {6, 10}. So the count is 2

// With dynamic programming, the complexity is O(sum*n)
// without dynamic programming, the complexity would be O(2^n)

func countSetsEqualingSum(numbers [] int, sum int) int {
  results := make(map[string]int)
  return countSets(numbers, sum, len(numbers)-1, results)
}

func countSets(numbers [] int, sum, index int, results map[string]int) int {
  key := strconv.Itoa(sum) + ":" + strconv.Itoa(index)
  if value, ok := results[key]; ok {
    return value
  }
  
  var result int
  if sum == 0 {
    return 1
  } else if sum < 1 {
    return 0
  } else if index < 0 {
    return 0
  } else if sum < numbers[index] {
    result = countSets(numbers, sum, index-1, results)
  } else {
    first := countSets(numbers, sum-numbers[index], index-1, results)
    second := countSets(numbers, sum, index-1, results)
    result = first + second
  }
  
  results[key] = result
  return result
}

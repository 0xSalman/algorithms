package array

import "sort"

// Problem: given an integer array, find all pairs whose
// sum equals M where M >= 1
// assume array is not sorted and may contain duplicates

// This problem could be solved in O(n) but requires additional space.
// In Go, the extra space is tiny due to use of an empty struct.
// However, when the array is too large and space cost
// is not desirable then this could be solved in O(nlogn) with sort

type pair struct {
  first, second int
}

func findPairsFaster(numbers []int, sum int) []pair {
  if len(numbers) <= 1 {
    return []pair{}
  }
  
  numbersMap := make(map[int]struct{})
  pairs := make([]pair, 0)
  for _, num := range numbers {
    target := sum - num
    if _, ok := numbersMap[target]; ok {
      pairs = append(pairs, pair{num, target})
    } else {
      numbersMap[num] = struct{}{}
    }
  }
  
  return pairs
}

func findPairsEfficiently(numbers []int, sum int) []pair {
  if len(numbers) <= 1 {
    return []pair{}
  }
  
  sort.Ints(numbers)
  pairs := make([]pair, 0)
  right := len(numbers) - 1
  for left := 0; left < right; {
    target := numbers[left] + numbers[right]
    if target == sum {
      pairs = append(pairs, pair{numbers[left], numbers[right]})
      left++
      right--
    } else if target < sum {
      left++
    } else {
      right--
    }
  }
  
  return pairs
}

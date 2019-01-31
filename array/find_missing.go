package array

// Problem: given an integer array [1...n], find all missing numbers
// between the range [a, b]. The numbers in the array are a<=number<=b
// assume array is not sorted and may contain duplicates

func findMissingNumbers(numbers []int, start, end int) []int {
  missing := make([]int, 0)
  count := len(numbers)
  
  if count <= 1 {
    return missing
  }
  
  numbersMap := make(map[int]struct{})
  for _, num := range numbers {
    numbersMap[num] = struct{}{}
  }
  
  for i := start; i <= end; i++ {
    if _, ok := numbersMap[i]; ok == false {
      missing = append(missing, i)
    }
  }
  
  return missing
}

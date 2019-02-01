package puzzle

import "testing"

func TestFindMinimumNumberOfCameras_Simple(t *testing.T) {
  spaces := []int{1, 2, 3, 4, 5}
  expected := 2
  actual := findMinimumNumberOfCameras(1, spaces)
  if expected != actual {
    t.Errorf("Got %d, wanted %d\n", actual, expected)
  }
}

func TestFindMinimumNumberOfCameras_SmallRange(t *testing.T) {
  spaces := []int{2, 4, 5, 6, 7, 9, 11, 12}
  expected := 3
  actual := findMinimumNumberOfCameras(2, spaces)
  if expected != actual {
    t.Errorf("Got %d, wanted %d\n", actual, expected)
  }
}

func TestFindMinimumNumberOfCameras_LargeRange(t *testing.T) {
  spaces := []int{1, 15, 30, 40, 50}
  expected := 3
  actual := findMinimumNumberOfCameras(10, spaces)
  if expected != actual {
    t.Errorf("Got %d, wanted %d\n", actual, expected)
  }
}

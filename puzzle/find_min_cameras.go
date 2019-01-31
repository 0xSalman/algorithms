package puzzle

// You are given a task of installing video cameras that will detect where cars are parked
// on one side of the street. Each camera can cover one or more parking spaces around it,
// as defined by cameraRange parameter. Camera will always be installed on top of a specific parking space.

// Your job is to implement a function that, given a list of all possible parking space locations
// on the street, will return a minimum number of cameras needed to detect cars in all those parking
// spaces. Each parking space spans 1 unit of space, and spaces are numbered starting at 1 - so no
// elements in parkingSpaces array will be less than 1. parkingSpaces array will always be sorted
// from low to high, as in examples.

// REMEMBER that camera can only be installed on top of an existing parking space.

// Example 1:
// cameraRange = 1
// parkingSpaces = [1 , 2, 3, 4, 5]
//
// The result here is 2. If you place one camera on space 2, it will cover spaces 1, 2,
// and 3 (covers both ways), and camera on space 4 will cover space 4 and 5. It's ok
// to "double-cover" spaces, and note that we ARE looking at parkingSpaces[0].

func findMinimumNumberOfCameras(cameraRange int, parkingSpace [] int) int {
  var numberOfCameras int
  var startSpace int
  var endSpace int
  // var cameraSpot int
  
  for i, spot := range parkingSpace {
    start := spot
    if withinRange(len(parkingSpace), i-1) &&
      spot-cameraRange <= parkingSpace[i-1] {
      start = spot - cameraRange
    }
    
    end := spot
    if withinRange(len(parkingSpace), i+1) &&
      spot+cameraRange >= parkingSpace[i+1] {
      end = spot + cameraRange
    }
    
    if start == startSpace {
      endSpace = end
      // cameraSpot = spot
    } else if spot > endSpace {
      numberOfCameras++
      // cameraSpot = spot
      startSpace = spot
      endSpace = end
    }
    
    // fmt.Printf("(start, end): (%d, %d)\n", startSpace, endSpace)
    // fmt.Printf("(spot, numCameras, cameraSpot): (%d, %d, %d)\n", spot, numberOfCameras, cameraSpot)
  }
  
  return numberOfCameras
}

func withinRange(length, index int) bool {
  return index >= 0 && index < length
}

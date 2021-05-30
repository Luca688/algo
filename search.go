package algo

import "math"

func BinarySearch(slice []int, target int) (res bool, index int) {
	length := len(slice)

	if length < 1 || target > slice[length-1] || target < slice[0] {
		return false, 0
	}

	min, max := 0, length-1

	for {
		if slice[max] == target {
			return true, max
		}
		if slice[min] == target {
			return true, min
		}

		if max-min <= 1 {
			return false, 0
		}

		mid := int(math.Floor(float64((max + min) / 2)))
		if target > slice[mid] {
			min = mid
		} else {
			max = mid
		}

	}

}

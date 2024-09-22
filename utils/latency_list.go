package utils

// FindMin is a function that finds the minimum value in a list of float64 values

func FindMin(arr []float64) float64 {
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

// FindMax is a function that finds the maximum value in a list of float64 values

func FindMax(arr []float64) float64 {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

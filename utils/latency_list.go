package utils

func FindMin(arr []float64) float64 {
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func FindMax(arr []float64) float64 {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

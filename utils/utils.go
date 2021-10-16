package utils

func MinFromSlice(slice []int) int {
	smallestNumber := slice[0]
	for _, element := range slice {
		if element < smallestNumber {
			smallestNumber = element
		}
	}
	return smallestNumber
}

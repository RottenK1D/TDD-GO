package array

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sumSlice []int

	for _, number := range slices {
		sumSlice = append(sumSlice, Sum(number))
	}

	return sumSlice
}

func SumAllTails(slices ...[]int) []int {
	var sumSlice []int

	for _, number := range slices {
		tail := number[1:]
		sumSlice = append(sumSlice, Sum(tail))
	}
	return sumSlice
}

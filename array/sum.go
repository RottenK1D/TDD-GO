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
		if len(number) == 0 {
			sumSlice = append(sumSlice, 0)
		} else {
			tail := number[1:]
			sumSlice = append(sumSlice, Sum(tail))
		}
	}
	return sumSlice
}

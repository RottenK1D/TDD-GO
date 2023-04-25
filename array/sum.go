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

	for i, x := range slices[0] {
		for _, slice := range slices[1:] {
			if i >= len(slice) {
				break
			}

			x += slice[i]
		}

		sumSlice = append(sumSlice, x)
	}

	return sumSlice
}

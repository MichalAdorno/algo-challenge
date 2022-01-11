package arrays

func MaxWater(A []int) int {
	if A == nil || len(A) < 3 {
		return 0
	}
	return lowToHighAndEqual(A) + highToLow(A)
}

func lowToHighAndEqual(A []int) int {
	i := 0
	j := 1
	N := len(A)
	acc := 0
	delta := 0

	for i < N && j < N {
		if A[i] > A[j] {
			delta += (A[i] - A[j])
			j += 1
		} else if A[i] <= A[j] {
			acc += delta
			delta = 0
			i = j
			j = i + 1
		}
	}
	return acc
}

func highToLow(A []int) int {
	N := len(A)
	i := N - 1
	j := N - 2
	acc := 0
	delta := 0

	for i >= 0 && j >= 0 {
		//fmt.Printf("1. i=%d, j=%d\n", i, j)
		if A[i] >= A[j] {
			delta += (A[i] - A[j])
			j -= 1
			//fmt.Printf("2/ i=%d, j=%d\n", i, j)
		} else if A[i] <= A[j] {
			acc += delta
			delta = 0
			i = j
			j = i - 1
			//fmt.Printf("3/ i=%d, j=%d\n", i, j)
		}
	}
	return acc
}

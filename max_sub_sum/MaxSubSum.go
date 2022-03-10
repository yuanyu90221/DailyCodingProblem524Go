package max_sub_sum

func MaxSubSum(nums []int) int {
	max := 0
	accum := 0
	for _, val := range nums {
		accum += val
		if accum < val { // Sn < A[i]
			accum = val
		}
		if max < accum { // max < Sn
			max = accum
		}
	}
	return max
}

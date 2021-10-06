package slicex

func DistinctInt32(nums []int32) []int32 {
	if len(nums) <= 1 {
		return nums
	}
	distinctMap := make(map[int32]struct{}, len(nums))
	for _, n := range nums {
		distinctMap[n] = struct{}{}
	}
	distincNums := nums[0:0]
	for n := range distinctMap {
		distincNums = append(distincNums, n)
	}
	return distincNums
}

package code

func ContainsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		ii := i + k
		for j := i + 1; j <= ii; j++ {
			if j >= len(nums) {
				break
			}
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}

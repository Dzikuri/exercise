package binarysearch

func BinarySearch(nums []int, target int) int {

	result := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = i
		}
	}

	return result
}

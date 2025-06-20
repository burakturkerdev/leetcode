package leetcode912

func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			hold := nums[0]
			nums[0] = nums[1]
			nums[1] = hold
		}
		return nums
	}

	mid := len(nums) / 2

	return putIntoPlace(sortArray(nums[:mid]), sortArray(nums[mid:]))
}

// I know, it's a terrible function name.
func putIntoPlace(nums1 []int, nums2 []int) []int {
	merged := make([]int, len(nums1)+len(nums2))

	p1 := 0
	p2 := 0
	i := 0
	for p1 != len(nums1) && p2 != len(nums2) {
		if nums1[p1] < nums2[p2] {
			merged[i] = nums1[p1]
			p1++
		} else {
			merged[i] = nums2[p2]
			p2++
		}
		i++
	}
	for p1 < len(nums1) {
		merged[i] = nums1[p1]
		p1++
		i++
	}

	for p2 < len(nums2) {
		merged[i] = nums2[p2]
		p2++
		i++
	}

	return merged
}

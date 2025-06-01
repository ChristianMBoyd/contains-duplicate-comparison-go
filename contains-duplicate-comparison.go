package containsDuplicate

import (
	"slices"
)

// UsingMap implements set-like behavior in Go via a map[int]bool.
// This is probably the most obvious O(N) hash map solution.
func UsingMap(nums []int) bool {
	previousNums := make(map[int]bool)
	for _, value := range nums {
		if previousNums[value] {
			return true
		}
		previousNums[value] = true
	}
	return false
}

// LeetCodeSolution is the 0ms solution copy-pasted from the
// contains-duplicate LeetCode problem submissions.
func LeetCodeSolution(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		isSorted := true
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] == nums[j+1] {
				return true
			}

			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
			}
		}

		if isSorted {
			break
		}
	}

	return false
}

// LeetCodeSolutionRefactored is LeetCodeSolution refactored to use
// modern Go and cleaned up for readability.  The purpose of keeping
// the copy-pasted form is to validate that these are the same.
func LeetCodeSolutionRefactored(nums []int) bool {
	// Todo - actually refactor, this is just range syntax at the moment.
	for i := range len(nums) - 1 {
		isSorted := true
		for j := range len(nums) - 1 - i {
			if nums[j] == nums[j+1] {
				return true
			}

			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
			}
		}

		if isSorted {
			break
		}
	}

	return false
}

// UsingSort calls Go's slices.Sort to do an in-place sort on the
// inputs before checking for duplicates.  The expected time
// complexity is O(NlogN) due to slices.Sort -- would be shocking if
// this is not O(NlogN) -- followed by a linear search.
func UsingSort(nums []int) bool {
	slices.Sort(nums)
	for i := range len(nums) - 1 {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

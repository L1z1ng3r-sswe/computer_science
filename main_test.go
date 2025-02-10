package main

import "testing"

func FindTarget(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}

func BinaryFindTarget(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

func BenchmarkFindTarget(b *testing.B) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	target := 19

	for i := 0; i < b.N; i++ {
		FindTarget(nums, target)
	}
}

func BenchmarkBinaryFindTarget(b *testing.B) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	target := 19

	for i := 0; i < b.N; i++ {
		BinaryFindTarget(nums, target)
	}
}

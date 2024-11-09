#include <iostream>
#include <vector>

void bubbleSort(std::vector<int>& nums) {
	size_t n = nums.size();
	for (size_t i = 0; i < n - 1; ++i) {
		for (size_t j = 0; j < n-i-1; ++j) {
			if (nums[j] > nums[j+1]){
				std::swap(nums[j], nums[j+1]);
			}
		}
	}
}
from typing import List


class Solution:
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        start, end = 0, len(nums) - 1
        while start < end:
            while start < end:
                if nums[start] % 2 == 1:
                    break
                start += 1
            while start < end:
                if nums[end] % 2 == 0:
                    break
                end -= 1
            if start < end:
                nums[start], nums[end] = nums[end], nums[start]
                start, end = start + 1, end - 1
        return nums

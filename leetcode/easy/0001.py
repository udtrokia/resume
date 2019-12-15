class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_set = set(nums)
        for i, j in enumerate(nums):
            if (target - j) in num_set:
                i2 = nums.index(target-j)
                if i != i2:
                    return [i, i2]

from typing import List
from sortedcontainers import SortedList

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        n = len(nums)

        x = SortedList()
        x.add((nums[0], 1))
        result = 1
        for i in range(1, n):
            v = nums[i]
            z = 1
            j = x.bisect_right((v-1, n+1))
            if j > 0:
                z = x[j-1][1] + 1
            
            # since the lengths are stricly increasing
            # if we need to pop an element it has to be x[i]
            if j < len(x):
                if x[j][1] <= z:
                    x.pop(j)

            x.add((v, z))
            print(x)
            if z > result:
                result = z

        return result

def main():
    sol = Solution()

    nums = [10,9,2,5,3,7,101,18]
    print(sol.lengthOfLIS(nums))

    nums = [0,1,0,3,2,3]
    print(sol.lengthOfLIS(nums))

    nums = [7,7,7,7,7,7,7]
    print(sol.lengthOfLIS(nums))

if __name__ == "__main__":
    main()

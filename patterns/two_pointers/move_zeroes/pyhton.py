class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i=0
        for j in range(0,len(nums),1 ):
            if (i==j) or  (nums[i]==0 and nums[j]==0):
                continue
            if (nums[i]==0) and (nums[j]!=0):
                nums[i], nums[j]=nums[j],nums[i]
            i+=1



# | |
# 101010101

# ||
# 01030
#  | |
# 10030
#   | |
# 13000
#   | |
# 13000
#  | |
# 100039

# if i==j ||  nums[i]==0 && nums[j]==0{
#     continue
# }
# if nums[i]==0 && nums[j]!=0{
#    nums[i], nums[j]=nums[j],nums[i]
#    i++
# }


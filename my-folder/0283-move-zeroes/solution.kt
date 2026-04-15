class Solution {
    fun moveZeroes(nums: IntArray): Unit {
        val temporary = mutableListOf<Int>()

        for (n in nums) {
            if (n != 0) {
                temporary.add(n)
            }
        }

        for (i in nums.indices) {
            if (i < temporary.size) {
                nums[i] = temporary[i]
            } else {
                nums[i] = 0
            }
        }
    }
}

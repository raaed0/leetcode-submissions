function numberGame(nums: number[]): number[] {
    nums.sort((a, b) => a-b)
    let l = nums.length - 1;
    for (let i = 0; i < l; i += 2) {
        [nums[i], nums[i+1]] = [nums[i+1], nums[i]]
    }

    return nums
};

# kandane algo
# find the subString until its sum reaches to zero and maintain the var to keep track the max sum value
func maxSubArray(nums []int) int {
    maxSum, curSum := math.MinInt64, 0
    for i, _ := range nums{
        curSum += nums[i]
        if curSum > maxSum {
            maxSum = curSum
        }
        if curSum < 0 {
            curSum = 0
        }
    }
    return maxSum
}

    

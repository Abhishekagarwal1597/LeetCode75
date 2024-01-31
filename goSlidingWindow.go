#Go solution usign sliding window
func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    
    for i:=0;i<k;i++{
        sum += nums[i]
    }

    maxSum := sum
    
    for j:=k;j<len(nums);j++{
        sum += nums[j] - nums[j-k]
        if sum > maxSum {
            maxSum = sum
        }
    }
    return float64(maxSum)/float64(k)

}


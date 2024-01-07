#using 2 pointer approach
#Maximum operation required to get the K-Sum

func maxOperations(nums []int, k int) int {
    // sort the list
    sort.Ints(nums)
    operation := 0
    // left := 0
    // right := len(nums)-1 
    for left,right := 0,len(nums)-1;left<right;{
    // for left < right {
        sum := nums[left] + nums[right]
        if sum == k {
            operation++
            left++
            right--
        }else if sum > k {
            right--
        }else {
            left++
        }
    }
    return operation    
}


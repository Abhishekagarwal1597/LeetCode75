#Use 2 Pointer approach one for iterating through array and one for moving non-zero element
func moveZeroes(nums []int)  {
    var index int
    for _, v := range nums {
        if v != 0 {
            nums[index] = v
            index++
        }
    }
    
    for index < len(nums) {
        nums[index] = 0
        index++    
    }
    
}

or Use the below method,
func moveZeroes(nums []int) {
    // Use two pointers: one for iterating through the array and another for placing non-zero elements.
    nonZeroIndex := 0

    // Iterate through the array
    for i := 0; i < len(nums); i++ {
        // If the current element is non-zero, swap it with the non-zero index
        if nums[i] != 0 {
            nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
            nonZeroIndex++
        }
    }
}



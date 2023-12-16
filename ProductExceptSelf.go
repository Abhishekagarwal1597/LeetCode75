#use 2 pointer apporch to resolve this
#find the product of left and right side of array index
#combine them as needed
timeComplixity = o(n)

func productExceptSelf(nums []int) []int {
    n := len(nums)
    
    leftMultiplication := make([]int, n)
    rightMultiplication := make([]int, n)
    productArray := make([]int, n)
  
    leftMul := 1
    rightMul := 1
    
    for i:=0; i < n ; i++{
        leftMultiplication[i] = leftMul
        leftMul *=  nums[i]    
    }
    
    for i:=n-1; i >= 0; i--{
        rightMultiplication[i] = rightMul
        rightMul *=   nums[i]
    }
    
    // var productArray int
    var result []int
    
    for i := range nums{
       productArray[i] =  leftMultiplication[i] * rightMultiplication[i]
        result = append(result, productArray[i])
    }
    
    return productArray
}
